import { Sidebar, Badge } from '@puppet/react-components';
import * as React from 'react';

type MainLayoutProps = {
  children: React.ReactNode;
};

export const MainLayout = ({ children }: MainLayoutProps) => {
  return (
    <div className="app-view">
      <Sidebar>
        <Sidebar.Header logo="pipelines" ariaLabel="Return to the home page" as="a" href="/" />
        <Sidebar.Navigation>
          <Sidebar.Item onClick={() => {}} title="Home" icon="home" active containerElement="div" />
          <Sidebar.Section label="reports">
            <Sidebar.Item onClick={() => {}} title="Code" icon="code" />
            <Sidebar.Item onClick={() => {}} title="Build" icon="build" count={5} />
            <Sidebar.Item
              onClick={() => {}}
              title="Deploy"
              icon="rocket"
              badge={
                <Badge pill type="success">
                  4
                </Badge>
              }
            />
          </Sidebar.Section>
          <Sidebar.Section label="config">
            <Sidebar.Item
              onClick={() => {}}
              title="Connections"
              icon="connections"
              label="config"
            />
          </Sidebar.Section>
        </Sidebar.Navigation>
        <Sidebar.Footer
          profileIcon={
            <img src="https://www.gravatar.com/avatar/205e460b479e2e5b48aec07710c08d50?s=100" />
          }
          username="Lorem Ipsum"
          version="1969.7.20"
          onClick={console.log}
          enableSignout
          onSignout={console.log}
          signoutTooltip="This is a custom tooltip"
        />
      </Sidebar>
      <div className="app-content no-margins">
        <main className="flex-1 relative overflow-y-auto focus:outline-none">{children}</main>
      </div>
    </div>
  );
};
