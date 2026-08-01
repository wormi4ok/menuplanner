<template>
  <div>
    <WeekSlider v-if="isMobile" :menu="data" />
    <WeekGrid v-else :menu="data" />
    <b-loading :is-full-page="true" v-model="isLoading" />
  </div>
</template>

<script>
import WeekGrid from '@/components/WeekGrid.vue';
import WeekSlider from '@/components/WeekSlider.vue';
import isMobile from '@/isMobile';
import { mapActions, mapState } from 'pinia';
import { useWeekStore } from '@/stores/week';
import { useRecipesStore } from '@/stores/recipes';
import { useCoursesStore } from '@/stores/courses';

export default {
  name: 'Home',
  components: {
    WeekSlider,
    WeekGrid,
  },
  data: () => ({
    isLoading: false,
  }),
  computed: {
    ...mapState(useWeekStore, {
      data: 'weekMenu',
    }),
    isMobile: () => isMobile.value,
  },
  async mounted() {
    this.isLoading = true;
    try {
      await Promise.all([
        this.fetchCurrentWeek(),
        this.fetchRecipes(),
        this.fetchCourses(),
      ]);
    } finally {
      this.isLoading = false;
    }
  },
  methods: {
    ...mapActions(useWeekStore, ['fetchCurrentWeek']),
    ...mapActions(useRecipesStore, ['fetchRecipes']),
    ...mapActions(useCoursesStore, ['fetchCourses']),
  },
};
</script>
