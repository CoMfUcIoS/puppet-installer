import { Sidebar, Badge } from '@puppet/react-components';
import * as React from 'react';

type MainLayoutProps = {
  children: React.ReactNode;
};

export const MainLayout = ({ children }: MainLayoutProps) => {
  return (
    <div className="app-view">
      <Sidebar>
        <Sidebar.Header logo="enterprise" ariaLabel="Return to the home page" as="a" href="/" />
        <Sidebar.Navigation>
          <Sidebar.Item onClick={() => {}} title="Deploy" icon="rocket" />
          <Sidebar.Item onClick={() => {}} title="Upgrade" icon="download" />
          <Sidebar.Item onClick={() => {}} title="Migrate" icon="import" />
          <Sidebar.Item onClick={() => {}} title="Convert" icon="build" />
          <Sidebar.Section label="Management">
            <Sidebar.Item onClick={() => {}} title="Add component" icon="module" />
            <Sidebar.Item onClick={() => {}} title="Backup" icon="folder" />
            <Sidebar.Item
              onClick={() => {}}
              title="Restore"
              icon="refresh"
              badge={
                <Badge pill type="success">
                  4
                </Badge>
              }
            />
            <Sidebar.Item onClick={() => {}} title="Status" icon="combo" />
          </Sidebar.Section>
        </Sidebar.Navigation>
        <Sidebar.Footer username="Lorem Ipsum" version="0.0.1" />
      </Sidebar>
      <div className="app-content no-margins">
        <main className="flex-1 relative overflow-y-auto focus:outline-none">{children}</main>
      </div>
    </div>
  );
};
