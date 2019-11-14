<!--
  - Developed by Richard Malaschitz on 4/1/19 11:37 AM
  - Last modified 4/1/19 11:17 AM
  - Copyright (c) 2019. All right reserved.
  -
  -->

<template>
  <v-container fluid>
    <v-toolbar flat color="white">
      <v-icon large color="teal darken-2">home</v-icon>
      <v-toolbar-title>
        {{ person.FirstName }} {{ person.Surname }} ({{ person.DtoBirthDate }})
        <v-btn flat small color="primary" :to="'/person/' + person.ID">detail</v-btn>
      </v-toolbar-title>
      <v-spacer />
    </v-toolbar>

    <v-form v-model="valid" ref="form" lazy-validation>
      <v-container>
        <v-layout wrap>
          <v-flex xs6 md3 sm4>
            <v-menu
              v-model="menuDatePicker1"
              :close-on-content-click="false"
              :nudge-right="40"
              lazy
              transition="scale-transition"
              offset-y
              full-width
              min-width="290px"
            >
              <template v-slot:activator="{ on }">
                <v-text-field
                  v-model="visit.DtoDatum.Date"
                  label="Dátum návštevy"
                  prepend-icon="event"
                  readonly
                  v-on="on"
                  required
                  clearable
                />
              </template>
              <v-date-picker v-model="visit.DtoDatum.Date" @input="menuDatePicker1 = false" />
            </v-menu>
          </v-flex>

          <v-flex xs6 md3 sm4>
            <v-menu
              v-model="menuTimePicker1"
              :close-on-content-click="false"
              :nudge-right="40"
              lazy
              transition="scale-transition"
              offset-y
              full-width
              min-width="290px"
            >
              <template v-slot:activator="{ on }">
                <v-text-field
                  v-model="visit.DtoDatum.Time"
                  label="Čas návštevy"
                  prepend-icon="access_time"
                  readonly
                  v-on="on"
                  required
                  clearable
                />
              </template>
              <v-time-picker
                v-model="visit.DtoDatum.Time"
                @input="menuTimePicker1 = false"
                format="24hr"
              />
            </v-menu>
          </v-flex>
          <v-flex xs6 md3 sm4>
            <v-text-field
              v-model="visit.Duration"
              label="Trvanie v minútach"
              type="number"
              prepend-icon="timer"
              required
            />
          </v-flex>
          <v-flex xs6 md3 sm4>
            <v-select
              :items="cars"
              label="Auto"
              v-model="visit.CarID"
              item-text="Name"
              item-value="ID"
              clearable
            ></v-select>
          </v-flex>
        </v-layout>

        <v-layout wrap>
          <v-flex xs6 md3 sm4>
            <v-menu
              v-model="menuDatePicker2"
              :close-on-content-click="false"
              :nudge-right="40"
              lazy
              transition="scale-transition"
              offset-y
              full-width
              min-width="290px"
            >
              <template v-slot:activator="{ on }">
                <v-text-field
                  v-model="visit.DtoVyjazdFrom.Date"
                  label="Dátum výjazdu"
                  prepend-icon="event"
                  readonly
                  v-on="on"
                  required
                  clearable
                />
              </template>
              <v-date-picker v-model="visit.DtoVyjazdFrom.Date" @input="menuDatePicker2 = false" />
            </v-menu>
          </v-flex>

          <v-flex xs6 md3 sm4>
            <v-menu
              v-model="menuTimePicker2"
              :close-on-content-click="false"
              :nudge-right="40"
              lazy
              transition="scale-transition"
              offset-y
              full-width
              min-width="290px"
            >
              <template v-slot:activator="{ on }">
                <v-text-field
                  v-model="visit.DtoVyjazdFrom.Time"
                  label="Čas výjazdu"
                  prepend-icon="access_time"
                  readonly
                  v-on="on"
                  required
                  clearable
                />
              </template>
              <v-time-picker
                v-if="menuTimePicker2"
                v-model="visit.DtoVyjazdFrom.Time"
                format="24hr"
                full-width
                @click:minute="menuTimePicker2 = false"
              />
            </v-menu>
          </v-flex>

          <v-flex xs6 md3 sm4>
            <v-menu
              v-model="menuDatePicker3"
              :close-on-content-click="false"
              :nudge-right="40"
              lazy
              transition="scale-transition"
              offset-y
              full-width
              min-width="290px"
            >
              <template v-slot:activator="{ on }">
                <v-text-field
                  v-model="visit.DtoVyjazdTo.Date"
                  label="Dátum príchodu"
                  prepend-icon="event"
                  readonly
                  v-on="on"
                  required
                  clearable
                />
              </template>
              <v-date-picker v-model="visit.DtoVyjazdTo.Date" @input="menuDatePicker3 = false" />
            </v-menu>
          </v-flex>

          <v-flex xs6 md3 sm4>
            <v-menu
              v-model="menuTimePicker3"
              :close-on-content-click="false"
              :nudge-right="40"
              lazy
              transition="scale-transition"
              offset-y
              full-width
              min-width="290px"
            >
              <template v-slot:activator="{ on }">
                <v-text-field
                  v-model="visit.DtoVyjazdTo.Time"
                  label="Čas príchodu"
                  prepend-icon="access_time"
                  readonly
                  v-on="on"
                  required
                  clearable
                />
              </template>
              <v-time-picker
                v-model="visit.DtoVyjazdTo.Time"
                v-if="menuTimePicker3"
                @click:minute="menuTimePicker3 = false"
                format="24hr"
              />
            </v-menu>
          </v-flex>
        </v-layout>

        <v-layout>
          <v-flex xs12>
            <v-autocomplete
              v-model="visit.Users"
              :items="users"
              chips
              label="Z Plamienka"
              item-text="Name"
              item-value="ID"
              multiple
            ></v-autocomplete>
          </v-flex>
          <v-flex xs12>
            <v-autocomplete
              v-model="visit.Persons"
              :items="person.DtoRelatives"
              chips
              label="Účastníci z domu - príbuzní"
              item-value="ID"
              multiple
            >
              <template v-slot:selection="{ item, selected }">
                <v-chip :selected="selected" color="blue-grey" class="white--text">
                  <span v-text="item.Person.FirstName + ' ' + item.Person.Surname"></span>
                </v-chip>
              </template>
              <template v-slot:item="{ item }">
                <v-list-tile-content>
                  <v-list-tile-title v-text="item.Relationship.Relation"></v-list-tile-title>
                  <v-list-tile-sub-title v-text="item.Person.FirstName + ' ' + item.Person.Surname"></v-list-tile-sub-title>
                </v-list-tile-content>
              </template>
            </v-autocomplete>
          </v-flex>
        </v-layout>
        <v-layout wrap>
          <v-flex xs3 md2 sm2>
            <h4>Typ návštevy:</h4>
          </v-flex>

          <v-flex xs6 md2 sm3>
            <v-checkbox v-model="visit.IsPlanned" label="Plánovaná"></v-checkbox>
          </v-flex>
          <v-flex xs6 md2 sm3>
            <v-checkbox v-model="visit.IsZdravotna" label="Zdravotná"></v-checkbox>
          </v-flex>
          <v-flex xs6 md2 sm3>
            <v-checkbox v-model="visit.IsSprevadzanie" label="Sprevádzanie"></v-checkbox>
          </v-flex>
          <v-flex xs6 md2 sm3>
            <v-checkbox v-model="visit.IsSocial" label="Sociálna"></v-checkbox>
          </v-flex>
          <v-flex xs6 md2 sm3>
            <v-checkbox v-model="visit.IsPoUmrti" label="Po úmrtí"></v-checkbox>
          </v-flex>
        </v-layout>

        <v-layout wrap>
          <v-flex xs12 md12 sm12>
            <v-text-field
              v-model="visit.Popis"
              label="Stručný popis"
              type="text"
              prepend-icon="create"
              required
            />
          </v-flex>
          <v-flex xs12>
            <v-textarea
              v-model="visit.PopisDetail"
              label="Popis"
              type="text"
              prepend-icon="create"
              required
            ></v-textarea>
          </v-flex>
          <v-flex xs12>
            <v-textarea
              v-model="visit.Vysetrenie"
              label="Vyšetrenie (zapíše sa do eZdravia)"
              type="text"
              prepend-icon="create"
              required
            ></v-textarea>
          </v-flex>
          <v-flex xs12>
            <v-textarea
              v-model="visit.MaterialnaPomoc"
              label="Materiálna pomoc"
              type="text"
              prepend-icon="create"
              required
            ></v-textarea>
          </v-flex>
          <v-flex xs12>
            <v-textarea
              v-model="visit.SocialnePoradenstvo"
              label="Sociálne poradenstvo"
              type="text"
              prepend-icon="create"
              required
            ></v-textarea>
          </v-flex>
        </v-layout>
        <v-layout>
          <v-btn color="info" @click="saveData">
            <v-icon left>save_alt</v-icon>Uložiť
          </v-btn>

          <v-btn color="warning" @click="readData">
            <v-icon left>refresh</v-icon>Refresh
          </v-btn>
        </v-layout>
      </v-container>
    </v-form>
  </v-container>
</template>

<script>
export default {
  name: 'VisitHome',

  props: ['personid', 'visitid'],

  data: () => ({
    visit: {
      DtoDatum: {
        Date: '',
        Time: ''
      },
      Users: [],
      Persons: [],
      Smer: true
    },
    person: {},
    users: [],
    cars: [],

    valid: false,
    menuDatePicker1: false,
    menuTimePicker1: false,
    menuDatePicker2: false,
    menuTimePicker2: false,
    menuDatePicker3: false,
    menuTimePicker3: false
  }),

  methods: {
    sendData: function () {},

    userName: function (id) {
      for (var i = 0; i < this.users.length; i++) {
        if (this.users[i].ID === id) {
          return this.users[i].Name
        }
      }
      return id
    },

    saveData: function () {
      if (!this.$refs.form.validate()) {
        return
      }
      this.$axios
        .post('/r/api/visithome', this.visit)
        .then(response => {
          if (response.status == 202) {
            this.$store.commit('alert', 'Chyba: ' + response.data.Error)
          } else {
            this.visit = response.data
            this.$store.commit('message', 'Uložené')
          }
        })
        .catch(response => {
          console.log('WRONG', response)
          this.$store.commit('alert', 'Chyba: ' + response)
        })
    },

    readData: function () {
      this.$axios
        .get('/r/api/person/' + this.personid)
        .then(response => {
          if (response.status === 202) {
            this.$store.commit('alert', 'Chyba: ' + response.data.Error)
          } else {
            this.person = response.data
          }
        })
        .catch(response => {
          this.$store.commit('alert', 'Chyba: ' + response)
        })

      this.$axios
        .get('/r/api/users')
        .then(response => {
          if (response.status == 202) {
            this.$store.commit('alert', 'Chyba: ' + response.data.Error)
          } else {
            this.users = response.data
          }
        })
        .catch(response => {
          this.$store.commit('alert', 'Chyba: ' + response)
        })

      this.$axios
        .get('/r/api/cars')
        .then(response => {
          if (response.status == 202) {
            this.$store.commit('alert', 'Chyba: ' + response.data.Error)
          } else {
            this.cars = response.data
          }
        })
        .catch(response => {
          this.$store.commit('alert', 'Chyba: ' + response)
        })

      if (this.visitid === 'new') {
        var today = new Date()
        this.visit = {
          DtoDatum: {
            Date:
              today.getFullYear() +
              '-' +
              String(today.getMonth() + 1).padStart(2, '0') +
              '-' +
              String(today.getDate()).padStart(2, '0'),
            Time: ''
          },
          Users: [],
          Persons: [],
          Smer: true
        }
        this.visit.Users.push(this.$store.state.id)
      } else {
        this.$axios
          .get('/r/api/visithome/' + this.visitid)
          .then(response => {
            console.log('OK', response)
            if (response.status == 202) {
              this.$store.commit('alert', 'Chyba: ' + response.data.Error)
            } else {
              this.visit = response.data
            }
          })
          .catch(response => {
            console.log('WRONG', response)
            this.$store.commit('alert', 'Chyba: ' + response)
          })
      }
    }
  },

  mounted () {
    this.$nextTick(() => {
      this.readData()
    })
  }
}
</script>

<style scoped></style>
