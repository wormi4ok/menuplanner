<template>
  <Week :menu="data"/>
</template>

<script>
import Week from '@/components/Week.vue';
import { mapGetters } from 'vuex';

export default {
  name: 'Home',
  components: {
    Week,
  },
  computed: mapGetters({
    data: 'weekMenu',
  }),
};
</script>
