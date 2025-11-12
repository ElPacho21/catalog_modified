# CatalogGo
Microservicio de Catalogo.

## Version: 1.0

**Contact information:**  
Nestor Marsollier  
nmarsollier@gmail.com  

---
### /articles

#### POST
##### Summary

Crear Artículo

##### Description

Crear Artículo

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| Authorization | header | Bearer {token} | Yes | string |
| body | body | Informacion del articulo | Yes | [article.UpdateArticleData](#articleupdatearticledata) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | Articulo | [article.ArticleData](#articlearticledata) |
| 400 | Bad Request | [errs.ValidationErr](#errsvalidationerr) |
| 401 | Unauthorized | [rst.ErrorData](#rsterrordata) |
| 404 | Not Found | [rst.ErrorData](#rsterrordata) |
| 500 | Internal Server Error | [rst.ErrorData](#rsterrordata) |

### /articles/:articleId

#### DELETE
##### Summary

Eliminar Artículo

##### Description

Eliminar Artículo

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| Authorization | header | Bearer {token} | Yes | string |
| articleId | path | ID de articlo | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | No Content |  |
| 400 | Bad Request | [errs.ValidationErr](#errsvalidationerr) |
| 401 | Unauthorized | [rst.ErrorData](#rsterrordata) |
| 404 | Not Found | [rst.ErrorData](#rsterrordata) |
| 500 | Internal Server Error | [rst.ErrorData](#rsterrordata) |

#### GET
##### Summary

Obtener un articulo

##### Description

Obtener un articulo

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| Authorization | header | Bearer {token} | Yes | string |
| articleId | path | ID de articlo | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | Articulo | [article.ArticleData](#articlearticledata) |
| 400 | Bad Request | [errs.ValidationErr](#errsvalidationerr) |
| 401 | Unauthorized | [rst.ErrorData](#rsterrordata) |
| 404 | Not Found | [rst.ErrorData](#rsterrordata) |
| 500 | Internal Server Error | [rst.ErrorData](#rsterrordata) |

#### POST
##### Summary

Actualizar Artículo

##### Description

Actualizar Artículo

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| Authorization | header | Bearer {token} | Yes | string |
| body | body | Informacion del articulo | Yes | [article.UpdateArticleData](#articleupdatearticledata) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | Articulo | [article.ArticleData](#articlearticledata) |
| 400 | Bad Request | [errs.ValidationErr](#errsvalidationerr) |
| 401 | Unauthorized | [rst.ErrorData](#rsterrordata) |
| 404 | Not Found | [rst.ErrorData](#rsterrordata) |
| 500 | Internal Server Error | [rst.ErrorData](#rsterrordata) |

### /articles/search/:criteria

#### GET
##### Summary

Obtener un articulo

##### Description

Obtener un articulo

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| Authorization | header | Bearer {token} | Yes | string |
| articleId | path | ID de articlo | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | Articulos | [ [article.ArticleData](#articlearticledata) ] |
| 400 | Bad Request | [errs.ValidationErr](#errsvalidationerr) |
| 401 | Unauthorized | [rst.ErrorData](#rsterrordata) |
| 404 | Not Found | [rst.ErrorData](#rsterrordata) |
| 500 | Internal Server Error | [rst.ErrorData](#rsterrordata) |

---
### /rabbit/article_exist

#### GET
##### Summary

Mensage Rabbit article_exist/article_exist

##### Description

Otros microservicios nos solicitan validar articulos en el catalogo.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| article_exist | body | Message para article_exist | Yes | [rschema.ConsumeArticleExist](#rschemaconsumearticleexist) |

##### Responses

| Code | Description |
| ---- | ----------- |

### /rabbit/logout

#### GET
##### Summary

Mensage Rabbit logout

##### Description

Escucha de mensajes logout desde auth.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| body | body | Estructura general del mensage | Yes | [rbt.InputMessage-string](#rbtinputmessage-string) |

##### Responses

| Code | Description |
| ---- | ----------- |

### /rabbit/order_placed

#### GET
##### Summary

Mensage Rabbit order_placed/order_placed

##### Description

Cuando se recibe el mensage order_placed damos de baja al stock para reservar los articulos. Queda pendiente enviar mensaje confirmando la operacion al MS de Orders.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| order_placed | body | Message order_placed | Yes | [rschema.ConsumeOrderPlacedMessage](#rschemaconsumeorderplacedmessage) |

##### Responses

| Code | Description |
| ---- | ----------- |

---
### Models

#### article.ArticleData

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| _id | string |  | No |
| description | string |  | No |
| enabled | boolean |  | No |
| image | string |  | No |
| name | string |  | No |
| price | number |  | No |
| stock | integer |  | No |

#### article.UpdateArticleData

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| description | string |  | Yes |
| image | string |  | No |
| name | string |  | Yes |
| price | number |  | Yes |
| stock | integer |  | Yes |

#### errs.ValidationErr

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| messages | [ [errs.errField](#errserrfield) ] |  | No |

#### errs.errField

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| message | string |  | No |
| path | string |  | No |

#### rbt.InputMessage-string

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| correlation_id | string | *Example:* `"123123"` | No |
| exchange | string | *Example:* `"Remote Exchange to Reply"` | No |
| message | string |  | No |
| routing_key | string | *Example:* `"Remote RoutingKey to Reply"` | No |

#### rschema.ConsumeArticleExist

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| correlation_id | string | *Example:* `"123123"` | No |
| exchange | string | *Example:* `"Remote Exchange to Reply"` | No |
| message | [rschema.ConsumeArticleExistMessage](#rschemaconsumearticleexistmessage) |  | No |
| routing_key | string | *Example:* `"Remote RoutingKey to Reply"` | No |

#### rschema.ConsumeArticleExistMessage

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| articleId | string | *Example:* `"ArticleId"` | No |
| referenceId | string | *Example:* `"Remote Reference Object Id"` | No |

#### rschema.ConsumeOrderPlacedArticle

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| articleId | string |  | No |
| quantity | integer |  | No |

#### rschema.ConsumeOrderPlacedMessage

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| articles | [ [rschema.ConsumeOrderPlacedArticle](#rschemaconsumeorderplacedarticle) ] |  | No |
| cartId | string |  | No |
| orderId | string |  | No |

#### rst.ErrorData

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| error | string |  | No |
