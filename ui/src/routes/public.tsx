import { Suspense } from 'react';
import { Navigate, Outlet } from 'react-router-dom';

import { MainLayout } from '@/components/Layout';
import { Landing } from '@/features/misc';

const App = () => {
  return (
    <MainLayout>
      <Suspense
        fallback={<div className="h-full w-full flex items-center justify-center">Loading....</div>}
      >
        <Outlet />
      </Suspense>
    </MainLayout>
  );
};

export const publicRoutes = [
  {
    path: '/',
    element: <App />,
    children: [
      { path: '/', element: <Landing /> },
      { path: '*', element: <Navigate to="." /> },
    ],
  },
];
