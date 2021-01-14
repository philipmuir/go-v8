#include "v8golang.h"
#include "v8.h"

extern "C" {

void v8_Initialize() {
//  const char* flags = "--expose_gc";
//  v8::V8::SetFlagsFromString(flags, strlen(flags));

  auto platform_ = v8::platform::NewDefaultPlatform();
  v8::V8::InitializePlatform(platform_.get());

  v8::V8::Initialize();
  return;
}

const char* Version() {
  return v8::V8::GetVersion();
}

  Version version = {V8_MAJOR_VERSION, V8_MINOR_VERSION, V8_BUILD_NUMBER, V8_PATCH_LEVEL};

}

//  IsolatePtr v8_Isolate_New(StartupData startupData) {
//    v8::Isolate::CreateParams create_params;
//    create_params.array_buffer_allocator = allocator;
//
//    if (startupData.length > 0 && startupData.data != NULL) {
//      v8::StartupData* data = new v8::StartupData;
//      data->data = startupData.data;
//      data->raw_size = startupData.length;
//      create_params.snapshot_blob = data;
//    }else {
//      // not needed but lets be explicit about this (v8 source code: src/api/api.cc:8524)
//      create_params.snapshot_blob = nullptr;
//    }
//
//    v8::Isolate* isolate_ptr = v8::Isolate::New(create_params);
//    return static_cast<IsolatePtr>(isolate_ptr);
//  }
//
//  void v8_Isolate_Terminate(IsolatePtr isolate_ptr) {
//    v8::Isolate* isolate = static_cast<v8::Isolate*>(isolate_ptr);
//    isolate->TerminateExecution();
//  }

//
//void v8_Isolate_RequestGarbageCollectionForTesting(IsolatePtr pIsolate) {
//  ISOLATE_SCOPE(static_cast<v8::Isolate*>(pIsolate));
//
//  isolate->RequestGarbageCollectionForTesting(v8::Isolate::kFullGarbageCollection);
//}
//
//HeapStatistics v8_Isolate_GetHeapStatistics(IsolatePtr pIsolate) {
//  if (pIsolate == NULL) {
//    return HeapStatistics{0};
//  }
//
//  ISOLATE_SCOPE(static_cast<v8::Isolate*>(pIsolate));
//
//  v8::HeapStatistics hs;
//  isolate->GetHeapStatistics(&hs);
//
//  return HeapStatistics{
//    hs.total_heap_size(),
//    hs.total_heap_size_executable(),
//    hs.total_physical_size(),
//    hs.total_available_size(),
//    hs.used_heap_size(),
//    hs.heap_size_limit(),
//    hs.malloced_memory(),
//    hs.peak_malloced_memory(),
//    hs.does_zap_garbage()
//  };
//}
//
//void v8_Isolate_LowMemoryNotification(IsolatePtr pIsolate) {
//  if (pIsolate == NULL) {
//    return;
//  }
//  ISOLATE_SCOPE(static_cast<v8::Isolate*>(pIsolate));
//  isolate->LowMemoryNotification();
//}
//
//void v8_Isolate_Release(IsolatePtr isolate_ptr) {
//  if (isolate_ptr == nullptr) {
//    return;
//  }
//  v8::Isolate* isolate = static_cast<v8::Isolate*>(isolate_ptr);
//  isolate->Dispose();
//}