import { mapState } from 'pinia';
import { useCoursesStore } from '@/stores/courses';
import { useWeekStore } from '@/stores/week';

const courseBySlot = {
  0: 'breakfast',
  1: 'main',
  2: 'main',
};

export const weekDays = ['Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday', 'Sunday'];

export const currentWeekDay = () => (new Date().getDay() || 7) - 1;

export default {
  data: () => ({
    weekDays,
    slots: Object.keys(courseBySlot).map(Number),
  }),
  computed: {
    ...mapState(useCoursesStore, [
      'listCourses',
    ]),
  },
  methods: {
    course(slot) {
      return this.listCourses.find((course) => course.name === courseBySlot[slot]);
    },
    removeSlot(day, slot) {
      useWeekStore().emptySlot({ day, slot });
    },
    fillSlot({ day, slot }, recipe) {
      useWeekStore().fillSlot({ day, slot, recipe });
    },
  },
};
