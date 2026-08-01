<template>
  <b-tabs
    class="recipes"
    v-model="selectedDay"
    v-touch:swipe.left="onSwipeLeft"
    v-touch:swipe.right="onSwipeRight"
    expanded
    type="is-toggle-rounded"
  >
    <b-tab-item v-for="(day, i) in menu" :key="i" :label="weekDayLabels[i]">
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
import WeekSlots, { currentWeekDay } from '@/mixins/WeekSlots';

export default {
  name: 'WeekSlider',
  components: {
    MenuSlot,
    DailySummary,
  },
  props: {
    menu: Object,
  },
  data: () => ({
    weekDayLabels: ['Mo', 'Tu', 'We', 'Th', 'Fr', 'Sa', 'Su'],
    selectedDay: currentWeekDay(),
  }),
  methods: {
    onSwipeLeft() {
      if (this.selectedDay >= this.weekDays.length - 1) {
        return;
      }
      this.selectedDay += 1;
      this.announceDay();
    },
    onSwipeRight() {
      if (this.selectedDay <= 0) {
        return;
      }
      this.selectedDay -= 1;
      this.announceDay();
    },
    announceDay() {
      this.$buefy.toast.open({
        duration: 1300,
        message: this.weekDays[this.selectedDay],
        position: 'is-bottom',
        type: 'is-light',
      });
    },
  },
  mixins: [
    WeekSlots,
  ],
};
</script>

<style>
.b-tabs.recipes .is-toggle-rounded a {
  padding: 0.5em 0.5em;
}
</style>
