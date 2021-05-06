<template>
  <b-tabs expanded type="is-toggle-rounded">
    <b-tab-item v-for="(day, i) in menu" :key="i" :label="weekDays[i]">
      <MenuSlot
        v-for="(recipe, slot) in day.recipes"
        :key="''.concat('slider',slot,day)"
        :recipe="recipe"
        :course="course(slot)"
        class="block"
        @delete-recipe="removeSlot(i,slot)"
        @pick-recipe="fillSlot({day:i, slot}, $event)"
      />
      <div class="section">
        <h5 class="title is-5">Summary</h5>
        <DailySummary :recipes="day.recipes"/>
      </div>
    </b-tab-item>
  </b-tabs>
</template>

<script>

import MenuSlot from '@/components/MenuSlot.vue';
import DailySummary from '@/components/DailySummary.vue';
import { mapGetters } from 'vuex';

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
      weekDays: ['Mo', 'Tu', 'We', 'Th', 'Fr', 'Sa', 'Su'],
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

</style>
