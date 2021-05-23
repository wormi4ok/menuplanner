<template>
  <div>
    <WeekSlider v-if="this.$root.isMobile" :menu="data"/>
    <WeekGrid v-else :menu="data"/>
    <b-loading :is-full-page="true" v-model="isLoading"></b-loading>
  </div>
</template>

<script>
import WeekGrid from '@/components/WeekGrid.vue';
import WeekSlider from '@/components/WeekSlider.vue';
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
  computed: mapGetters({
    data: 'weekMenu',
  }),
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
