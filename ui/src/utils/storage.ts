const storagePrefix = 'installation_';

const storage = {
  getInstallation: (name: string) => {
    return JSON.parse(window.localStorage.getItem(`${storagePrefix}${name}`) as string);
  },
  setInstallation: (name: string) => {
    window.localStorage.setItem(`${storagePrefix}token`, JSON.stringify(name));
  },
  clearInstallation: (name: string) => {
    window.localStorage.removeItem(`${storagePrefix}${name}`);
  },
};

export default storage;
