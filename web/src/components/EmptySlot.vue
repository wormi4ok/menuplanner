<template>
  <div class="card">
    <div class="box">
      <b-field label="Pick a recipe">
        <b-autocomplete
          v-model="name"
          :data="filteredDataObj"
          field="name"
          placeholder="Recipe name"
          rounded
          clearable
          @select="pickRecipe">
          <template #empty>No recipes found</template>
        </b-autocomplete>
      </b-field>
    </div>
  </div>
</template>

<script>
import { mapGetters } from 'vuex';

export default {
  name: 'EmptySlot',
  data: () => ({
    name: '',
  }),
  computed: {
    filteredDataObj() {
      return this.data.filter((option) => (
        option.name
          .toString()
          .toLowerCase()
          .indexOf(this.name.toLowerCase()) >= 0
      ));
    },
    ...mapGetters({
      data: 'listRecipes',
    }),
  },
  methods: {
    pickRecipe(option) {
      this.$emit('pick-recipe', option);
    },
  },
};
</script>

<style scoped>

</style>
