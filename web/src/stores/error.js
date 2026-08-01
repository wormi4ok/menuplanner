import { defineStore } from 'pinia';

export const useErrorStore = defineStore('error', {
  state: () => ({
    error: '',
  }),
  getters: {
    getError: (state) => state.error,
  },
  actions: {
    reportError(errorMsg) {
      this.error = errorMsg;
    },
  },
});

export default useErrorStore;
