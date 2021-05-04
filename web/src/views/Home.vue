<template>
  <WeekGrid :menu="data"/>
</template>

<script>
import WeekGrid from '@/components/WeekGrid.vue';
import { mapGetters } from 'vuex';

export default {
  name: 'Home',
  components: {
    WeekGrid,
  },
  computed: mapGetters({
    data: 'weekMenu',
  }),
};
</script>
