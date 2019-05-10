/* Code generated by cmd/cgo; DO NOT EDIT. */

/* package storj.io/storj/lib/uplink/ext */


#line 1 "cgo-builtin-export-prolog"

#include <stddef.h> /* for ptrdiff_t below */

#ifndef GO_CGO_EXPORT_PROLOGUE_H
#define GO_CGO_EXPORT_PROLOGUE_H

#ifndef GO_CGO_GOSTRING_TYPEDEF
typedef struct { const char *p; ptrdiff_t n; } _GoString_;
#endif

#endif

/* Start of preamble from import "C" comments.  */


#line 6 "apikey.go"

 #include <stdlib.h>
 #ifndef UPLINK_HEADERS
   #define UPLINK_HEADERS
   #include "headers/main.h"
 #endif

#line 1 "cgo-generated-wrapper"

#line 6 "common.go"

 #include <stdlib.h>
 #ifndef UPLINK_HEADERS
   #define UPLINK_HEADERS
   #include "headers/main.h"
 #endif

#line 1 "cgo-generated-wrapper"

#line 6 "project.go"

 #ifndef UPLINK_HEADERS
   #define UPLINK_HEADERS
   #include "headers/main.h"
 #endif

#line 1 "cgo-generated-wrapper"

#line 6 "uplink.go"

 #ifndef UPLINK_HEADERS
   #define UPLINK_HEADERS
   #include "headers/main.h"
 #endif

#line 1 "cgo-generated-wrapper"


/* End of preamble from import "C" comments.  */


/* Start of boilerplate cgo prologue.  */
#line 1 "cgo-gcc-export-header-prolog"

#ifndef GO_CGO_PROLOGUE_H
#define GO_CGO_PROLOGUE_H

typedef signed char GoInt8;
typedef unsigned char GoUint8;
typedef short GoInt16;
typedef unsigned short GoUint16;
typedef int GoInt32;
typedef unsigned int GoUint32;
typedef long long GoInt64;
typedef unsigned long long GoUint64;
typedef GoInt64 GoInt;
typedef GoUint64 GoUint;
typedef __SIZE_TYPE__ GoUintptr;
typedef float GoFloat32;
typedef double GoFloat64;
typedef float _Complex GoComplex64;
typedef double _Complex GoComplex128;

/*
  static assertion to make sure the file is being used on architecture
  at least with matching size of GoInt.
*/
typedef char _check_for_64_bit_pointer_matching_GoInt[sizeof(void*)==64/8 ? 1:-1];

#ifndef GO_CGO_GOSTRING_TYPEDEF
typedef _GoString_ GoString;
#endif
typedef void *GoMap;
typedef void *GoChan;
typedef struct { void *t; void *v; } GoInterface;
typedef struct { void *data; GoInt len; GoInt cap; } GoSlice;

#endif

/* End of boilerplate cgo prologue.  */

#ifdef __cplusplus
extern "C" {
#endif


// ParseAPIKey parses an API Key

extern APIKey ParseAPIKey(char* p0, char** p1);

// Serialize serializes the API Key to a string

extern char* Serialize(APIKey p0);

extern struct IDVersion GetIDVersion(unsigned int p0, char** p1);

extern struct Bucket CreateBucket(GoUintptr p0, char* p1, struct BucketConfig p2, char** p3);

extern struct Uplink NewUplink(struct Config p0, char** p1);

extern Project OpenProject(struct Uplink p0, char* p1, APIKey p2, struct ProjectOptions p3, char** p4);

#ifdef __cplusplus
}
#endif
