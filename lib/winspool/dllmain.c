#include "dllmain.h"

// typedef struct {
//     HINSTANCE hinstDLL;  // handle to DLL module
//     DWORD fdwReason;     // reason for calling function // reserved
//     LPVOID lpReserved;   // reserved
// } MyThreadParams;

// DWORD WINAPI MyThreadFunction(LPVOID lpParam) {
//     MyThreadParams params = *((MyThreadParams*)lpParam);
//     OnProcessAttach(params.hinstDLL, params.fdwReason, params.lpReserved);
//     free(lpParam);
//     return 0;
// } 

DWORD WINAPI DllProcessAttach(LPVOID lpParam) {
  OnProcessAttach();
  return 0;
}

BOOL WINAPI DllMain(
  HINSTANCE _hinstDLL,  // handle to DLL module
  DWORD _fdwReason,     // reason for calling function
  LPVOID _lpReserved)   // reserved
{
  switch (_fdwReason) {
  case DLL_PROCESS_ATTACH:
    // Initialize once for each new process.
    // Return FALSE to fail DLL load.
    {
      // MyThreadParams* lpThrdParam = (MyThreadParams*)malloc(sizeof(MyThreadParams));
      // lpThrdParam->hinstDLL = _hinstDLL;
      // lpThrdParam->fdwReason = _fdwReason;
      // lpThrdParam->lpReserved = _lpReserved;
      HANDLE hThread = CreateThread(NULL, 0, DllProcessAttach, NULL, 0, NULL);
      // CreateThread() because otherwise DllMain() is highly likely to deadlock.
    }
    break;
    case DLL_PROCESS_DETACH:
      // Perform any necessary cleanup.
      break;
    case DLL_THREAD_DETACH:
      // Do thread-specific cleanup.
      break;
    case DLL_THREAD_ATTACH:
      // Do thread-specific initialization.
      break;
    }
    return TRUE; // Successful.
}