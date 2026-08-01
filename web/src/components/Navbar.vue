<template>
  <b-navbar class="mb-2">
    <template #brand>

      <b-navbar-item tag="router-link" :to="{ name: 'Home' }">
        <img src="@/assets/flaningo.jpeg" alt="MenuPlanner logo" />
      </b-navbar-item>

    </template>
    <template #start>

      <b-navbar-item class="is-tab" tag="router-link" :to="{ name: 'Home' }">
        Week
      </b-navbar-item>

      <b-navbar-item class="is-tab" tag="router-link" :to="{ name: 'Recipes' }">
        Recipes
      </b-navbar-item>
    </template>
    <template #end>

      <b-navbar-item tag="div" v-if="$route.name === 'Recipes'">
        <b-button class="is-primary" label="Add Recipe" @click="onAddRecipe" />
      </b-navbar-item>

      <b-navbar-item tag="div" class="mr-5" v-else>

        <b-button v-if="hasGaps" label="Fill gaps" class="is-primary" @click="onFillGaps" />
        <b-button v-else label="Clear week" class="is-danger" @click="onClearWeek" />

      </b-navbar-item>
      <NavbarUserMenu />
    </template>
  </b-navbar>
  <b-modal v-model="showAddRecipeForm" has-modal-card>
    <AddRecipeForm @close="showAddRecipeForm = false" />
  </b-modal>
</template>

<script>
import AddRecipeForm from '@/components/AddRecipeForm.vue';
import NavbarUserMenu from '@/components/NavbarUserMenu.vue';
import { mapActions, mapState } from 'pinia';
import { useWeekStore } from '@/stores/week';

export default {
  name: 'Navbar',
  components: {
    AddRecipeForm,
    NavbarUserMenu,
  },
  data: () => ({
    showAddRecipeForm: false,
  }),
  computed: {
    ...mapState(useWeekStore, [
      'hasGaps',
    ]),
  },
  methods: {
    ...mapActions(useWeekStore, ['fillGaps', 'emptyWeek']),
    onFillGaps() {
      this.fillGaps();
    },
    onClearWeek() {
      this.$buefy.dialog.confirm({
        message: 'Remove all recipes chosen for the week?',
        onConfirm: () => this.emptyWeek(),
      });
    },
    onAddRecipe() {
      this.showAddRecipeForm = true;
    },
  },
};
</script>

<style scoped>

</style>
