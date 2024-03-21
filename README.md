## POST /create/agreement
Ejemplo que debemos enviar al body:
>{
  "counterparty": {
    "IDRefCTPY": "123456789",
    "EmailCTPY": "counterparty@example.com",
    "LinkShareCTPY": "https://example.com/counterparty"
  },
  "agreementText": {
    "Text": "Este es el texto del acuerdo..."
  },
  "counterparty_signature": {
    "DNI": "12345678X",
    "FullName": "Firma del Contraparte",
    "Accepte": true
  },
  "proposing_firm": {
    "DNI": "87654321Y",
    "FullName": "Firma de la Propuesta",
    "Accepte": true
  },
  "agreement_status": {
    "Status": "active",
    "Text": "El contrato está activo"
  },
  "contractLinkId": {
    "ID": 123456,
    "Password": "contraseña_secreta"
  },
  "add_field_required":[
    {
    "field":"comprobante",
    "type":"img"
    },
    {
    "field":"comprobante2",
    "type":"img"
    }
  ]
}
