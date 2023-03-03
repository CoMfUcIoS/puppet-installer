import * as React from 'react';

type MainLayoutProps = {
  children: React.ReactNode;
};


export const MainLayout = ({ children }: MainLayoutProps) => {
  return (
      <div className="h-screen flex overflow-hidden bg-gray-100">
        <div className="flex flex-col w-0 flex-1 overflow-hidden">
          <main className="flex-1 relative overflow-y-auto focus:outline-none">{children}</main>
        </div>
      </div>
  );
};