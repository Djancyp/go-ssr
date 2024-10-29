export {}; // Ensure this file is treated as a module

declare global {
  interface globalThis {
    props: {
      name: string;
    };
  }
}
