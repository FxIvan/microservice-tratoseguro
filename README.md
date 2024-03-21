## POST /auth/signin
Ejemplo que debemos enviar al body:
>{
  "username": "fxivanTESTConEmailTt",
  "email":"jeffersonlinux666@gmail.com",
  "password": "contraseña123"
}

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

Necesitamos enviarle el token ya que la ruta esta proteguida
![image](https://github.com/FxIvan/microservice-tratoseguro/assets/62405720/a93f3ce7-28b9-4345-9862-cf48aa47a3fc)

