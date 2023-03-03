import * as React from 'react';
import { ErrorBoundary } from 'react-error-boundary';
import { HelmetProvider } from 'react-helmet-async';
import { QueryClientProvider } from 'react-query';
import { ReactQueryDevtools } from 'react-query/devtools';
import { BrowserRouter as Router } from 'react-router-dom';

import { Notifications } from '@/components/Notifications/Notifications';
import { queryClient } from '@/lib/react-query';

const ErrorFallback = () => {
  return (
    <div
      className=""
      role="alert"
    >
      <h2 className="">Ooops, something went wrong :( </h2>
      <div className="" onClick={() => window.location.assign(window.location.origin)}>
        Refresh
      </div>
    </div>
  );
};

type AppProviderProps = {
  children: React.ReactNode;
};

export const AppProvider = ({ children }: AppProviderProps) => {
  return (
    <React.Suspense
      fallback={
        <div className="flex items-center justify-center w-screen h-screen">
          Loading....
        </div>
      }
    >
      <ErrorBoundary FallbackComponent={ErrorFallback}>
        <HelmetProvider>
          <QueryClientProvider client={queryClient}>
            {process.env.NODE_ENV !== 'test' && <ReactQueryDevtools />}
            <Notifications />

              <Router>{children}</Router>

          </QueryClientProvider>
        </HelmetProvider>
      </ErrorBoundary>
    </React.Suspense>
  );
};