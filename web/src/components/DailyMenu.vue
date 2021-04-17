<template>
  <div class="menu">
    <MenuSlot :recipe="recipes[0]"
              @delete-recipe="deleteRecipe(0)"
              @pick-recipe="pickRecipe(0, $event)"/>
    <MenuSlot :recipe="recipes[1]"
              @delete-recipe="deleteRecipe(1)"
              @pick-recipe="pickRecipe(1, $event)"/>
    <MenuSlot :recipe="recipes[2]"
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
  name: 'Weekday',
  components: {
    MenuSlot,
  },
  props: {
    recipes: Object,
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
