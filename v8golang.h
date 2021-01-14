#ifndef MAINN_H
#define MAINN_H
#ifdef __cplusplus
extern "C" {
#endif

//#include <stdlib.h>
//#include <string.h>
#include <stddef.h>
//#include <stdint.h>
//#include <stdbool.h>

typedef struct {
    size_t totalHeapSize;
    size_t totalHeapSizeExecutable;
    size_t totalPhysicalSize;
    size_t totalAvailableSize;
    size_t usedHeapSize;
    size_t heapSizeLimit;
    size_t mallocedMemory;
    size_t peakMallocedMemory;
    size_t doesZapGarbage;
} HeapStatistics;
typedef struct {
    const char* data;
    int length;
} String;

typedef String Error;
typedef String StartupData;
typedef String ByteArray;

typedef void* IsolatePtr;
typedef void* ContextPtr;
typedef void* ValuePtr;
typedef void* PropertyDescriptorPtr;
typedef void* InspectorPtr;
typedef void* FunctionTemplatePtr;
typedef void* ObjectTemplatePtr;
typedef void* PrivatePtr;
typedef void* ExternalPtr;
typedef void* ResolverPtr;

extern void v8_Initialize();
extern IsolatePtr v8_Isolate_New(StartupData data);
extern void v8_Isolate_Terminate(IsolatePtr isolate);

extern const char* Version();

#ifdef __cplusplus
}
#endif


//extern void v8_Isolate_Release(IsolatePtr isolate);
//extern void v8_Isolate_RequestGarbageCollectionForTesting(IsolatePtr pIsolate);
//extern HeapStatistics v8_Isolate_GetHeapStatistics(IsolatePtr isolate);
//extern void v8_Isolate_LowMemoryNotification(IsolatePtr isolate);


#endif  // MAINN_H