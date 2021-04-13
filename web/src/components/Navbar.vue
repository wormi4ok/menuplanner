<template>
  <b-navbar>
    <template #brand>
      <b-navbar-item>
        <img src="https://flaningo.ru/wp-content/uploads/2019/03/cropped-99-1.jpg" alt="MenuPlanner logo">
      </b-navbar-item>
    </template>
    <template #start>
      <b-navbar-item href="#">
        <div class="buttons">
          <a class="button is-primary">
            <strong>Fill gaps</strong>
          </a>
        </div>
      </b-navbar-item>
      <b-navbar-item>
        <b-button label="Add Recipe" @click="showAddRecipeForm = true"/>
        <b-modal v-model="showAddRecipeForm">
          <AddRecipeForm  @close="showAddRecipeForm = false"/>
        </b-modal>
      </b-navbar-item>
    </template>
  </b-navbar>
</template>

<script>
import AddRecipeForm from '@/components/AddRecipeForm.vue';

export default {
  name: 'Navbar',
  components: {
    AddRecipeForm,
  },
  data() {
    return {
      showAddRecipeForm: false,
    };
  },
};
</script>

<style scoped>

</style>
