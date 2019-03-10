<template>
  <div>
    <v-btn icon @click="addTask()">
      <v-icon>add</v-icon>
    </v-btn>
    <div v-if="tasks.length > 0">
      <template v-for="(task, index) in tasks">
        <div :key="index">
          <v-layout row>  
            <v-flex xs1>
              <v-checkbox v-model="task.IsDeleted" label="" class="pl-5" @change="saveTask(task)" :disabled="task.IsDeleted" />
            </v-flex>
            <v-flex xs11>
              <v-textarea v-model="task.Ciel" label="" rows="1" auto-grow @change="saveTask(task)" :disabled="task.IsDeleted" />
            </v-flex>
          </v-layout>
        </div>
      </template>
    </div>
  </div>
</template>
<script>
export default {
  props: {
    personId: {
      type: String,
      required: true
    }
  },
  data: function() {
    return {
      meno: '',
      tasks: []
    }
  },
  computed: {},
  methods: {
    addTask: function() {
      let t = {}
      t.ID = null
      t.PersonID = this.personId
      t.Username = this.$store.state.name
      t.Deleted = null
      t.IsDeleted = false
      t.Ciel = ''
      this.tasks.push(t)
    },
    saveTask: function(task) {
      if (task.Ciel === '') {
        return false
      }
      var params = {
        ID: task.ID,
        IsDeleted: task.IsDeleted,
        Deleted: task.Deleted,
        Ciel: task.Ciel,
        PersonID: task.PersonID
      }
      this.$axios.post('/r/api/task', params)
        .then(response => {
          console.log("OK", response)
          if (response.status == 202) {
            this.$store.commit("alert", "Chyba: " + response.data.Error)
          } else {
            this.readData()
          }
        })
        .catch(response => {
          console.log("WRONG", response)
          this.$store.commit("alert", "Chyba: " + response)
        })
    },
    readData: function() {
      this.$axios.get('/r/api/tasks/' + this.personId)
        .then(response => {
          console.log("OK", response)
          if (response.status == 202) {
            this.$store.commit("alert", "Chyba: " + response.data.Error)
          } else {
            this.tasks = response.data
          }
        })
        .catch(response => {
          console.log("WRONG", response)
          this.$store.commit("alert", "Chyba: " + response)
        })
    }
  },
  mounted: function() {
    this.readData()
  }
}
</script>
