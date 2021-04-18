<template>
  <div class="menu">
    <MenuSlot :recipe="breakfast"
              @delete-recipe="deleteRecipe(0)"
              @pick-recipe="pickRecipe(0, $event)"/>
    <MenuSlot :recipe="lunch"
              @delete-recipe="deleteRecipe(1)"
              @pick-recipe="pickRecipe(1, $event)"/>
    <MenuSlot :recipe="dinner"
              @delete-recipe="deleteRecipe(2)"
              @pick-recipe="pickRecipe(2, $event)"/>
    <div class="block">
      Summary
    </div>
  </div>
</template>

<script>
import MenuSlot from '@/components/MenuSlot.vue';

export default {
  name: 'DailyMenu',
  components: {
    MenuSlot,
  },
  props: {
    recipes: Object,
  },
  data() {
    return {
      breakfast: this.recipes[0],
      lunch: this.recipes[1],
      dinner: this.recipes[2],
    };
  },
  methods: {
    deleteRecipe(slot) {
      this.$emit('empty-slot', slot);
    },
    pickRecipe(slot, recipe) {
      this.$emit('fill-slot', { slot, recipe });
    },
  },
};
</script>

<style scoped>

</style>
