<template>
  <div class="columns">
    <div :key="day" v-for="(today, day) in menu" class="column">
      <DailyMenu
        :recipes="today.recipes"
        @empty-slot="removeSlot(day,$event)"
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
    removeSlot(day, slot) {
      this.$store.dispatch('emptySlot', { day, slot });
    },
  },
};
</script>

<style scoped>

</style>
