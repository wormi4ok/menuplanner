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
import { mapActions, mapGetters } from 'vuex';

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
    ...mapGetters({
      data: 'weekMenu',
    }),
    isMobile: () => isMobile.value,
  },
  mounted() {
    this.isLoading = true;
    this.fetchCurrentWeek();
    this.fetchRecipes();
    this.fetchCourses();
    this.isLoading = false;
  },
  methods: {
    ...mapActions([
      'fetchCurrentWeek',
      'fetchRecipes',
      'fetchCourses',
    ]),
  },
};
</script>
