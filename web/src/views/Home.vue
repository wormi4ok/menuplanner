<template>
  <Week :menu="weekMenu"/>
</template>

<script>
import Week from '@/components/Week.vue';
import { mapActions, mapGetters } from 'vuex';

export default {
  name: 'Home',
  components: {
    Week,
  },
  computed: mapGetters(['weekMenu']),
  created() {
    this.fetchCurrentWeek();
  },
  methods: {
    ...mapActions(['fetchCurrentWeek']),
  },
};
</script>
