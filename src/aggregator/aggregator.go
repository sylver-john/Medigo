package aggregator

import(
  "log"
  m "../model"
)

func GetDrugOriginPercent(drugOrigins []m.Titulaire) map[string]float32 {
  percent := make(map[string]float32)
  for _, value := range drugOrigins {
    log.Println(value)
    for _, titulaire := range value.Titulaire {
      log.Println(titulaire)
      if _, ok := percent[titulaire]; ok {
        // si il existe déjà dans la liste des titulaires, on incrémente de 0.1
        percent[titulaire] += 0.1
      } else {
        // on ajoute, set à 0 et incrémente de 0.1
        percent[titulaire] = 0
        percent[titulaire] += 0.1
      }
    }
  }
  return percent
}
