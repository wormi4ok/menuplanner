<template>
  <div class="columns">
    <div :key="day" v-for="(today, day) in menu" class="column">
      <Menu :recipes="today.recipes"/>
    </div>
  </div>
</template>

<script>
import Menu from '@/components/Menu.vue';

export default {
  name: 'Week',
  components: {
    Menu,
  },
  props: {
    menu: Object,
  },
};
</script>

<style scoped>

</style>
