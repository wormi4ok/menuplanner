<template>
  <div class="columns">
    <div :key="day" v-for="(today, day) in menu" class="column">
      <DailyMenu
        :recipes="today.recipes"
        @sync-week="syncWeek"
      />
    </div>
  </div>
</template>

<script>
import DailyMenu from '@/components/DailyMenu.vue';

export default {
  name: 'Week',
  components: {
    DailyMenu,
  },
  props: {
    menu: Object,
  },
  methods: {
    syncWeek() {
      this.$emit('sync');
    },
  },
};
</script>

<style scoped>

</style>
