<template>
  <div class="grid">
    <div :key="day" v-for="day in weekDays" class="has-text-centered">
      <div class="subtitle">{{ day }}</div>
    </div>
    <template v-for="slot in courses">
      <MenuSlot
        v-for="(today, day) in menu"
        :key="''.concat(slot,day)"
        :recipe="today.recipes[slot]"
        :course="course(slot)"
        @delete-recipe="removeSlot(day,slot)"
        @pick-recipe="fillSlot({day, slot}, $event)"
      />
    </template>
    <DailySummary
      v-for="(today,day) in menu"
      :key="''.concat('summary',day)"
      :recipes="today.recipes"
    />
  </div>
</template>

<script>
import MenuSlot from '@/components/MenuSlot.vue';
import { mapGetters } from 'vuex';
import DailySummary from '@/components/DailySummary.vue';

export default {
  name: 'WeekGrid',
  components: {
    MenuSlot,
    DailySummary,
  },
  props: {
    menu: Object,
  },
  data() {
    return {
      weekDays: ['Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday', 'Sunday'],
      courses: [0, 1, 2],
    };
  },
  computed: {
    ...mapGetters([
      'listCourses',
    ]),
  },
  methods: {
    removeSlot(day, slot) {
      this.$store.dispatch('emptySlot', { day, slot });
    },
    fillSlot({ day, slot }, recipe) {
      this.$store.dispatch('fillSlot', { day, slot, recipe });
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

<style scoped>
.grid {
  display: grid;
  grid-auto-rows: min-content;
  grid-template-columns: repeat(7, 1fr );
  gap: 20px;
}
</style>
