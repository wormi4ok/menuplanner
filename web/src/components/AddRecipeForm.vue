<template>
  <div class="modal-card">
    <form @submit.prevent="AddRecipe">
      <header class="modal-card-head">
        <p class="modal-card-title">Add recipe</p>
      </header>
      <section class="modal-card-body">
        <b-field label="Name">
          <b-input type="text" required v-model="name"></b-input>
        </b-field>
        <b-field label="Description">
          <b-input type="text" v-model="description"></b-input>
        </b-field>
        <b-field label="Calories">
          <b-numberinput v-model="calories" step="50" min-step="1"></b-numberinput>
        </b-field>
        <b-field label="Protein">
          <b-numberinput v-model="protein" step="50" min-step="1"></b-numberinput>
        </b-field>
        <b-field label="Fat">
          <b-numberinput v-model="fat" step="50" min-step="1"></b-numberinput>
        </b-field>
        <b-field label="Carbs">
          <b-numberinput v-model="carbs" step="50" min-step="1"></b-numberinput>
        </b-field>
      </section>
      <footer class="modal-card-foot">
        <b-button label="Close" @click="$emit('close')"/>
        <b-button label="Create" class="is-primary" native-type="submit"/>
      </footer>
    </form>
  </div>
</template>

<script>
import api from '../api';

export default {
  name: 'AddRecipeForm',
  data() {
    return {
      name: '',
      description: '',
      imageUrl: '',
      calories: 0,
      protein: 0,
      fat: 0,
      carbs: 0,
    };
  },
  methods: {
    AddRecipe() {
      const recipe = {
        name: this.name,
        description: this.description,
        imageUrl: this.imageUrl,
        calories: this.calories,
        protein: this.protein,
        fat: this.fat,
        carbs: this.carbs,
      };
      api.recipe.create(recipe)
        .then(() => {
          this.$emit('close');
        })
        .catch((error) => {
          this.$emit('error', error.data);
        });
    },
  },
};
</script>

<style scoped>

</style>
