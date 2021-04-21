<template>
  <b-navbar class="mb-2">
    <template #brand>
      <b-navbar-item tag="router-link" :to="{ path: '/' }">
        <img src="https://flaningo.ru/wp-content/uploads/2019/03/cropped-99-1.jpg" alt="MenuPlanner logo">
      </b-navbar-item>
    </template>
    <template #start>
      <b-navbar-item>
        <a class="button is-primary" @click="onFillGaps">
          <strong>Fill gaps</strong>
        </a>
      </b-navbar-item>
      <b-navbar-item>
        <b-button label="Recipes" tag="router-link" :to="{ path: '/recipes' }"/>
      </b-navbar-item>
      <b-navbar-item>
        <b-button label="Add Recipe" @click="showAddRecipeForm = true"/>
        <b-modal v-model="showAddRecipeForm" scroll="keep">
          <AddRecipeForm @close="showAddRecipeForm = false"/>
        </b-modal>
      </b-navbar-item>
    </template>
  </b-navbar>
</template>

<script>
import AddRecipeForm from '@/components/AddRecipeForm.vue';
import { mapActions } from 'vuex';

export default {
  name: 'Navbar',
  components: {
    AddRecipeForm,
  },
  data: () => ({
    showAddRecipeForm: false,
  }),
  methods: {
    ...mapActions(['fillGaps']),
    onFillGaps() {
      this.fillGaps();
    },
  },
};
</script>

<style scoped>

</style>
