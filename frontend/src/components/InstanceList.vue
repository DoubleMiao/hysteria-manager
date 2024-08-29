<template>
  <div>
    <h2>Hysteria Instances</h2>
    <ul>
      <li v-for="instance in instances" :key="instance.id">
        {{ instance.name }} - {{ instance.port }}
        <button @click="editInstance(instance)">Edit</button>
        <button @click="deleteInstance(instance.id)">Delete</button>
      </li>
    </ul>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  data() {
    return {
      instances: []
    }
  },
  mounted() {
    this.fetchInstances();
  },
  methods: {
    fetchInstances() {
      axios.get('/instances')
        .then(response => {
          this.instances = response.data;
        })
        .catch(error => {
          console.error(error);
        })
    },
    editInstance(instance) {
      this.$router.push({ name: 'EditInstance', params: { id: instance.id } });
    },
    deleteInstance(id) {
      axios.delete(`/instances/${id}`)
        .then(() => {
          this.fetchInstances();
        })
        .catch(error => {
          console.error(error);
        });
    }
  }
}
</script>
