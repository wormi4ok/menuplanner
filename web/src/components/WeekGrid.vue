<template>
  <div class="week-grid">
    <div :key="day" v-for="(day, i) in weekDays" class="has-text-centered">
      <div class="subtitle" :class="{ 'has-text-weight-bold': today === i }">{{ day }}</div>
    </div>
    <template v-for="slot in slots">
      <MenuSlot
        v-for="(today, day) in menu"
        :key="''.concat(slot, day)"
        :recipe="today.recipes[slot]"
        :course="course(slot)"
        @delete-recipe="removeSlot(day, slot)"
        @pick-recipe="fillSlot({ day, slot }, $event)"
      />
    </template>
    <DailySummary
      v-for="(today, day) in menu"
      :key="''.concat('summary', day)"
      :recipes="today.recipes"
    />
  </div>
</template>

<script>
import MenuSlot from '@/components/MenuSlot.vue';
import DailySummary from '@/components/DailySummary.vue';
import WeekSlots, { currentWeekDay } from '@/mixins/WeekSlots';

export default {
  name: 'WeekGrid',
  components: {
    MenuSlot,
    DailySummary,
  },
  props: {
    menu: Object,
  },
  data: () => ({
    today: currentWeekDay(),
  }),
  mixins: [
    WeekSlots,
  ],
};
</script>

<style scoped>
.week-grid {
  display: grid;
  grid-auto-rows: min-content;
  grid-template-columns: repeat(7, 1fr );
  gap: 20px;
}
</style>
