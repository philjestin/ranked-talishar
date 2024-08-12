'use client';

import { Navigation } from "./Navigation";

export default function AppShell({ children }: { children: React.ReactNode }) {
  
  return (
    <>
      <div className="">
       <Navigation />
        <main className="bg-white">
          <div className="mx-auto max-w-7xl px-4 py-6 sm:px-6 lg:px-8 bg-white ">
            {children}
          </div>
        </main>
      </div>
    </>
  );
}
