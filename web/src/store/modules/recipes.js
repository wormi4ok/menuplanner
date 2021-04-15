// import axios from 'axios';

const state = {
  recipes: [
    {
      id: 1,
      name: 'Bolognese',
      calories: 600,
      protein: 40,
      fat: 50,
      carbs: 200,
    },
    {
      id: 2,
      name: 'Moroccan Carrot Soup',
      calories: 700,
      protein: 60,
      fat: 50,
      carbs: 50,
    },
    {
      id: 3,
      name: 'Pancake',
      calories: 450,
      protein: 50,
      fat: 10,
      carbs: 300,
    },
    {
      id: 4,
      name: 'Pasta Carbonara',
      calories: 300,
      protein: 120,
      fat: 80,
      carbs: 260,
    },
  ],
};

const getters = {
  listRecipes: (s) => s.recipes,
};

const actions = {};

const mutations = {};

export default {
  state,
  getters,
  actions,
  mutations,
};
