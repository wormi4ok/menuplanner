import { defineStore } from 'pinia';
import api from '@/api';
import { useErrorStore } from './error';

export const useCoursesStore = defineStore('courses', {
  state: () => ({
    courses: [],
  }),
  getters: {
    listCourses: (state) => state.courses,
  },
  actions: {
    async fetchCourses() {
      try {
        const response = await api.course.list();
        this.courses = response.data;
      } catch (error) {
        useErrorStore().reportError(error.response.data);
      }
    },
  },
});

export default useCoursesStore;
