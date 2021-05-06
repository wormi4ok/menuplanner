<template>
  <WeekSlider v-if="this.$root.isMobile" :menu="data"/>
  <WeekGrid v-else :menu="data"/>
</template>

<script>
import WeekGrid from '@/components/WeekGrid.vue';
import WeekSlider from '@/components/WeekSlider.vue';
import { mapGetters } from 'vuex';

export default {
  name: 'Home',
  components: {
    WeekSlider,
    WeekGrid,
  },
  computed: mapGetters({
    data: 'weekMenu',
  }),
};
</script>
