<template>
  <Week :menu="data"/>
</template>

<script>
import Week from '@/components/Week.vue';
import { mapActions, mapGetters } from 'vuex';

export default {
  name: 'Home',
  components: {
    Week,
  },
  computed: mapGetters({
    data: 'weekMenu',
  }),
  created() {
    this.fetchCurrentWeek();
    this.fetchRecipes();
  },
  methods: {
    ...mapActions([
      'fetchCurrentWeek',
      'fetchRecipes',
    ]),
  },
};
</script>
