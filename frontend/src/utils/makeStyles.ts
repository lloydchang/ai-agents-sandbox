import React, { Component, ReactNode, ErrorInfo } from 'react';

interface Props {
  children: ReactNode;
  fallback?: ReactNode;
}

interface State {
  hasError: boolean;
  error?: Error;
}

// Error boundary to catch Material-UI makeStyles errors
export class MaterialUIErrorBoundary extends Component<Props, State> {
  constructor(props: Props) {
    super(props);
    this.state = { hasError: false };
  }

  static getDerivedStateFromError(error: Error): State {
    return { hasError: true, error };
  }

  componentDidCatch(error: Error, errorInfo: ErrorInfo) {
    // Log the error but don't crash the app
    console.warn('Material-UI makeStyles error caught:', error.message);
    console.warn('Error info:', errorInfo);
  }

  render() {
    if (this.state.hasError) {
      // For the specific makeStyles error, render children anyway
      if (this.state.error?.message?.includes('refs')) {
        return this.props.children;
      }
      
      // For other errors, show fallback or error message
      return this.props.fallback || React.createElement('div', {
        style: {
          padding: 16,
          border: '1px solid #ff6b6b',
          borderRadius: 4,
          backgroundColor: '#ffe0e0'
        }
      }, [
        React.createElement('h3', { key: 'title' }, 'Something went wrong'),
        React.createElement('p', { key: 'message' }, this.state.error?.message)
      ]);
    }

    return this.props.children;
  }
}

export default MaterialUIErrorBoundary;
