<template>
  <v-container fluid>
    <v-toolbar flat color="white">
      <v-toolbar-title>Meniny {{ pondelok }} - {{ nedela }} </v-toolbar-title>
      <v-divider class="mx-2" inset vertical />
      <v-btn @click="minusDays(7)">
        Predchádzajúci týždeň
      </v-btn>
      <v-btn @click="addDays(7)">
        Nasledujúci týždeň
      </v-btn>
    </v-toolbar>

    <v-card
      v-for="m in meniny"
      :key="m.ID"
      class="ma-2 pa-2"
      :to="'/person/' + m.Person.ID"
    >
      <b>{{ m.Datum.Day }}.{{ m.Datum.Month }}.</b>
      {{ m.Typ }}
      <template v-if="m.Relationship.Sex">
        <span class="grey--text text--lighten-1"
          >{{ m.Person.FirstName }} {{ m.Person.Surname }}</span
        >
        -- <b>{{ m.Relationship.Relation }}</b> {{ m.Relative.FirstName }}
        {{ m.Relative.Surname }}
      </template>
      <template v-else>
        <b>{{ m.Person.FirstName }} {{ m.Person.Surname }}</b>
      </template>
    </v-card>
  </v-container>
</template>

<script>
export default {
  name: 'Meniny',

  data: () => ({
    today: new Date(),
    meniny: []
  }),

  mounted: function () {
    console.log('meniny')
    this.readData()
  },

  methods: {
    readData: function () {
      this.$axios
        .get('/r/api/meniny/' + this.pondelok)
        .then(response => {
          this.meniny = response.data
        })
        .catch(response => {
          console.log('WRONG', response)
          this.$store.commit('alert', 'Chyba: ' + response)
        })
    },

    addDays: function (d) {
      var newDate = new Date(this.today)
      newDate.setDate(newDate.getDate() + d)
      this.today = newDate
      this.readData()
    },

    minusDays: function (d) {
      var newDate = new Date(this.today)
      newDate.setDate(newDate.getDate() - d)
      this.today = newDate
      this.readData()
    }
  },

  computed: {
    pondelok: function () {
      var t = this.today.getDate() - ((this.today.getDay() + 6) % 7)
      var day = new Date(this.today)
      day.setDate(t)
      return day.getDate() + '.' + (day.getMonth() + 1) + '.'
    },

    nedela: function () {
      var t = this.today.getDate() - ((this.today.getDay() + 6) % 7) + 6
      var day = new Date(this.today)
      day.setDate(t)
      return day.getDate() + '.' + (day.getMonth() + 1) + '.'
    }
  }
}
</script>

<style scoped></style>
