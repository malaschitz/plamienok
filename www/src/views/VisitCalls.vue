<!--
  - Developed by Richard Malaschitz on 4/2/19 12:49 PM
  - Last modified 4/2/19 12:48 PM
  - Copyright (c) 2019. All right reserved.
  -
  -->

<template>
  <v-container fluid>
    <v-toolbar flat color="white">
      <v-toolbar-title>Telefonáty</v-toolbar-title>
    </v-toolbar>
    <v-data-table
      :headers="[
        { text: 'Dátum', value: 'DtoDatum' },
        { text: 'Typ', value: 'DtoTyp' },
        { text: 'Popis', value: 'Popis' },
        { text: 'Detailný Popis', value: 'PopisDetail' },
        { text: 'Akcie', scrollable: false, value: 'ID' }
      ]"
      :items="visits"
      class="elevation-1"
    >
      <template v-slot:items="props">
        <td class="text-left">
          {{ props.item.Datum.Day }}.{{ props.item.Datum.Month }}.{{
            props.item.Datum.Year
          }}
        </td>
        <td class="text-left">
          <v-chip
            v-if="props.item.Smer"
            color="primary"
            text-color="white"
            small
          >
            Z Plamienka
          </v-chip>
          <v-chip v-else color="primary" text-color="white" small>
            Z domu
          </v-chip>

          <v-tooltip bottom>
            <template v-slot:activator="{ on }">
              <v-chip
                v-if="props.item.IsZdravotna"
                color="info"
                text-color="white"
                small
              >
                Zdr
              </v-chip>
            </template>
            <span>Zdravotná</span>
          </v-tooltip>

          <v-tooltip bottom>
            <template v-slot:activator="{ on }">
              <v-chip
                v-if="props.item.IsSprevadzanie"
                color="info"
                text-color="white"
                small
              >
                Spr
              </v-chip>
            </template>
            <span>Sprevádzanie</span>
          </v-tooltip>

          <v-tooltip bottom>
            <template v-slot:activator="{ on }">
              <v-chip
                v-if="props.item.IsSocial"
                color="info"
                text-color="white"
                small
                v-on="on"
              >
                Soc
              </v-chip>
            </template>
            <span>Sociálna</span>
          </v-tooltip>

          <v-tooltip bottom>
            <template v-slot:activator="{ on }">
              <v-chip
                v-if="props.item.IsPoUmrti"
                color="info"
                text-color="white"
                small
              >
                Umr
              </v-chip>
            </template>
            <span>Po úmrtí</span>
          </v-tooltip>
        </td>
        <td class="text-left text-no-wrap">
          {{ props.item.Popis }}
        </td>
        <td
          class="text-left text-no-wrap text-truncate"
          style="max-width: 200px;"
        >
          {{ props.item.PopisDetail }}
        </td>
        <td>
          <v-icon class="mr-2" @click="editVisit(props.item)">
            edit
          </v-icon>
        </td>
      </template>
    </v-data-table>
  </v-container>
</template>

<script>
export default {
  name: 'VisitHomes',

  data: () => ({
    visits: []
  }),

  methods: {
    readData: function () {
      this.$axios
        .get('/r/api/visitcalls')
        .then(response => {
          console.log('OK', response)
          if (response.status == 202) {
            this.$store.commit('alert', 'Chyba: ' + response.data.Error)
          } else {
            this.visits = response.data
          }
        })
        .catch(response => {
          console.log('WRONG', response)
          this.$store.commit('alert', 'Chyba: ' + response)
        })
    },

    editVisit: function (item) {
      this.$router.push('/visitcall/' + item.PersonID + '/' + item.ID)
    }
  },

  mounted: function () {
    this.readData()
  },

  watch: {
    $route () {
      this.readData()
    }
  }
}
</script>

<style scoped></style>
