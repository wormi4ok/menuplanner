<template>
  <div class="columns">
    <div :key="day" v-for="(today, day) in menu" class="column">
      <DailyMenu :recipes="today.recipes"/>
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
};
</script>

<style scoped>

</style>
