import { defineStore } from 'pinia';
import api from '@/api';
import { useErrorStore } from './error';

function stateMerge(state, value, propName) {
  if (Object.prototype.toString.call(value) === '[object Object]'
    && (propName == null || Object.prototype.hasOwnProperty.call(state, propName))) {
    const o = propName == null ? state : state[propName];
    if (o != null) {
      Object.keys(value).forEach((prop) => stateMerge(o, value[prop], prop));
      return;
    }
  }
  state[propName] = value;
}

function emptyMenu() {
  return Object.fromEntries(
    Array.from({ length: 7 }, (_, day) => [day, { recipes: { 0: null, 1: null, 2: null } }]),
  );
}

export const useWeekStore = defineStore('week', {
  state: () => ({
    week: {
      menu: emptyMenu(),
    },
  }),
  getters: {
    weekMenu: (state) => state.week.menu,
    hasGaps: (state) => Object.values(state.week.menu).some(
      (day) => Object.values(day.recipes).some(
        (r) => !r || r.id === 0,
      ),
    ),
    recipePosition: (state) => (id) => {
      const position = [];
      Object.values(state.week.menu).some(
        (menu, day) => Object.values(menu.recipes).forEach(
          (recipe, slot) => {
            if (!recipe || recipe.id !== id) {
              return false;
            }
            position.push({ day, slot });
            return true;
          },
        ),
      );
      return position.length > 0 ? position : null;
    },
  },
  actions: {
    setCurrentWeek(week) {
      stateMerge(this.$state, week, 'week');
    },
    resetCurrentWeek() {
      Object.entries(this.week.menu).forEach((weekday) => {
        const [day, menu] = weekday;
        Object.keys(menu.recipes).forEach((slot) => {
          this.week.menu[day].recipes[slot] = null;
        });
      });
    },
    async fetchCurrentWeek() {
      try {
        const response = await api.week.getCurrent();
        this.setCurrentWeek(response.data);
      } catch (error) {
        useErrorStore().reportError(error.response.data);
      }
    },
    async emptySlot(config) {
      const currentWeek = this.week;
      currentWeek.menu[config.day].recipes[config.slot] = undefined;
      try {
        await api.week.delete(config.day, config.slot);
        this.setCurrentWeek(currentWeek);
      } catch (error) {
        useErrorStore().reportError(error.response.data);
      }
    },
    async fillSlot(config) {
      const currentWeek = this.week;
      currentWeek.menu[config.day].recipes[config.slot] = config.recipe;
      try {
        await api.week.update(currentWeek);
        this.setCurrentWeek(currentWeek);
      } catch (error) {
        useErrorStore().reportError(error.response.data);
      }
    },
    async fillGaps() {
      try {
        const response = await api.week.update(this.week, true);
        this.setCurrentWeek(response.data);
      } catch (error) {
        useErrorStore().reportError(error.response.data);
      }
    },
    emptyWeek() {
      Object.entries(this.week.menu).forEach((weekday) => {
        const [day, menu] = weekday;
        Object.keys(menu.recipes).forEach((slot) => {
          this.emptySlot({ day, slot });
        });
      });
    },
  },
});

export default useWeekStore;
