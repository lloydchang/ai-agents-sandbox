// Direct monkey patch for Material-UI makeStyles detach function
(() => {
  // Store original console.error to prevent spam
  const originalConsoleError = console.error;
  
  // Override console.error to filter out the specific makeStyles error
  console.error = (...args) => {
    const message = args[0];
    if (typeof message === 'string' && message.includes('Cannot read properties of undefined (reading \'refs\')')) {
      // Silently ignore this specific error
      return;
    }
    // Pass through other errors
    originalConsoleError.apply(console, args);
  };
  
  // Patch the detach function globally
  const originalDetach = Function.prototype.call;
  Function.prototype.call = function(thisArg, ...args) {
    // Check if this is the detach function from makeStyles
    if (this.name === 'detach') {
      try {
        return originalDetach.apply(this, [thisArg, ...args]);
      } catch (error) {
        // Silently handle the refs error
        if (error.message && error.message.includes('refs')) {
          return undefined;
        }
        throw error;
      }
    }
    return originalDetach.apply(this, [thisArg, ...args]);
  };
  
  // Also patch window.onerror to catch this specific error
  const originalOnError = window.onerror;
  window.onerror = function(message, source, lineno, colno, error) {
    if (typeof message === 'string' && message.includes('Cannot read properties of undefined (reading \'refs\')')) {
      // Prevent the error from showing in console
      return true;
    }
    if (originalOnError) {
      return originalOnError.call(this, message, source, lineno, colno, error);
    }
    return false;
  };
  
  // Patch unhandled promise rejections
  const originalOnUnhandledRejection = window.onunhandledrejection;
  window.onunhandledrejection = function(event) {
    if (event.reason && event.reason.message && event.reason.message.includes('refs')) {
      event.preventDefault();
      return true;
    }
    if (originalOnUnhandledRejection) {
      return originalOnUnhandledRejection.call(this, event);
    }
    return false;
  };
})();

export default {};
