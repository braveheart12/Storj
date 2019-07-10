// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package uplink_test

import (
	"context"
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"storj.io/storj/internal/testcontext"
	"storj.io/storj/internal/testplanet"
	"storj.io/storj/lib/uplink"
	"storj.io/storj/pkg/storj"
)

type expectedResult struct {
	path     string
	isPrefix bool
}

type challenge struct {
	commands        []string         // list commands that should result in the same set of responses
	expectedResults []expectedResult // results that should come back for each command above
}

type listTest struct {
	bucket     string      // bucket to upload to
	paths      []string    // test will create a file at every path here
	challenges []challenge // set of list commands and expected results
}

func TestList(t *testing.T) {
	testplanet.Run(t, testplanet.Config{
		SatelliteCount:   1,
		StorageNodeCount: 1,
		UplinkCount:      1},
		func(t *testing.T, ctx *testcontext.Context, planet *testplanet.Planet) {
			apiKey := planet.Uplinks[0].APIKey[planet.Satellites[0].ID()]
			satelliteAddr := planet.Satellites[0].Local().Address.Address

			tests := []listTest{
				{
					bucket: "bu1",
					paths: []string{
						"a-file",
						"a/b-file",
						"a/b/slash-file",
						"a/b////",
						"a/b//",
						"a/b//c-file",
						"a/b//c/",
						"//bob",
					},
					challenges: []challenge{
						// TODO(isaac): solve encryption problem, then test "//"
						{
							commands: []string{"a", "a/"},
							expectedResults: []expectedResult{
								{
									path:     "b/",
									isPrefix: true,
								},
								{
									path:     "b-file",
									isPrefix: false,
								},
							},
						},
						{
							commands: []string{"a/b", "a/b/"},
							expectedResults: []expectedResult{
								{
									path:     "/",
									isPrefix: true,
								},
								{
									path:     "slash-file",
									isPrefix: false,
								},
							},
						},
						{
							commands: []string{"a/b//"},
							expectedResults: []expectedResult{
								{
									path:     "c/",
									isPrefix: true,
								},
								{
									path:     "c-file",
									isPrefix: false,
								},
								{
									path:     "/",
									isPrefix: true,
								},
								{
									path:     "",
									isPrefix: false,
								},
							},
						},
						{
							commands: []string{"a/b///"},
							expectedResults: []expectedResult{
								{
									path:     "/",
									isPrefix: true,
								},
							},
						},
						{
							commands: []string{"a/b////"},
							expectedResults: []expectedResult{
								{
									path:     "",
									isPrefix: false,
								},
							},
						},
					},
				},
			}

			for _, test := range tests {
				runTest(ctx, t, apiKey, satelliteAddr, test)
			}
		})
}

func runTest(ctx context.Context, t *testing.T, apiKey, satelliteAddr string,
	test listTest) {

	cfg := &uplink.Config{}
	cfg.Volatile.TLS.SkipPeerCAWhitelist = true

	ul, err := uplink.NewUplink(ctx, cfg)
	require.NoError(t, err)
	defer ul.Close()

	key, err := uplink.ParseAPIKey(apiKey)
	require.NoError(t, err)

	p, err := ul.OpenProject(ctx, satelliteAddr, key)
	require.NoError(t, err)
	defer p.Close()

	_, err = p.CreateBucket(ctx, test.bucket, nil)
	require.NoError(t, err)

	encKey, err := p.SaltedKeyFromPassphrase(ctx, "my secret passphrase")
	require.NoError(t, err)

	// Make an encryption context
	access := uplink.NewEncryptionAccessWithDefaultKey(*encKey)

	bu, err := p.OpenBucket(ctx, test.bucket, access)
	require.NoError(t, err)

	// First upload files to all the specified paths
	for _, path := range test.paths {
		err = bu.UploadObject(ctx, path, strings.NewReader("hello world"), nil)
		require.NoError(t, err)
	}

	// Now run the challenges and check the results
	for _, challenge := range test.challenges {
		for _, command := range challenge.commands {
			results, err := bu.ListObjects(ctx, &uplink.ListOptions{
				Direction: storj.After,
				Cursor:    "",
				Prefix:    command,
				Recursive: false,
			})
			require.NoError(t, err)
			compareResults(t, results.Items, challenge.expectedResults)
		}
	}
}

func compareResults(t *testing.T, items []storj.Object, expected []expectedResult) {
	require.Equal(t, len(expected), len(items))
	sort.SliceStable(items, func(i, j int) bool { return items[i].Path < items[j].Path })
	sort.SliceStable(expected, func(i, j int) bool { return expected[i].path < expected[j].path })
	for i, item := range items {
		require.Equal(t, expected[i].path, item.Path)
		require.Equal(t, expected[i].isPrefix, item.IsPrefix)
	}
}
