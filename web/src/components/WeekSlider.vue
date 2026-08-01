<template>
  <b-tabs
    class="recipes"
    v-model="selectedDay"
    v-touch:swipe.left="onSwipeLeft"
    v-touch:swipe.right="onSwipeRight"
    expanded
    type="is-toggle-rounded"
  >
    <b-tab-item v-for="(day, i) in menu" :key="i" :label="weekDays[i]">
      <MenuSlot
        v-for="(recipe, slot) in day.recipes"
        :key="''.concat('slider', slot, day)"
        :recipe="recipe"
        :course="course(slot)"
        class="block"
        @delete-recipe="removeSlot(i, slot)"
        @pick-recipe="fillSlot({ day: i, slot }, $event)"
      />
      <div class="section">
        <h5 class="title is-5">Summary</h5>
        <DailySummary :recipes="day.recipes" />
      </div>
    </b-tab-item>
  </b-tabs>
</template>

<script>

import MenuSlot from '@/components/MenuSlot.vue';
import DailySummary from '@/components/DailySummary.vue';
import { mapState } from 'pinia';
import { useCoursesStore } from '@/stores/courses';
import { useWeekStore } from '@/stores/week';

export default {
  name: 'WeekSlider',
  components: {
    MenuSlot,
    DailySummary,
  },
  props: {
    menu: Object,
  },
  data() {
    return {
      weekDaysFull: ['Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday', 'Sunday'],
      weekDays: ['Mo', 'Tu', 'We', 'Th', 'Fr', 'Sa', 'Su'],
      selectedDay: (new Date().getDay() || 7) - 1,
    };
  },
  computed: {
    ...mapState(useCoursesStore, [
      'listCourses',
    ]),
  },
  methods: {
    removeSlot(day, slot) {
      useWeekStore().emptySlot({ day, slot });
    },
    fillSlot({ day, slot }, recipe) {
      useWeekStore().fillSlot({ day, slot, recipe });
    },
    onSwipeLeft() {
      if (this.selectedDay >= this.weekDaysFull.length - 1) {
        return;
      }
      this.selectedDay += 1;
      this.$buefy.toast.open({
        duration: 1300,
        message: this.weekDaysFull[this.selectedDay],
        position: 'is-bottom',
        type: 'is-light',
      });
    },
    onSwipeRight() {
      if (this.selectedDay <= 0) {
        return;
      }
      this.selectedDay -= 1;
      this.$buefy.toast.open({
        duration: 1300,
        message: this.weekDaysFull[this.selectedDay],
        position: 'is-bottom',
        type: 'is-light',
      });
    },
    course(slot) {
      const map = {
        0: 'breakfast',
        1: 'main',
        2: 'main',
        3: 'pudding',
      };
      return this.listCourses.find((course) => course.name === map[slot]);
    },
  },
};
</script>

<style>
.b-tabs.recipes .is-toggle-rounded a {
  padding: 0.5em 0.5em;
}
</style>
