# settlement
settlement POC

## à faire

À Faire (pas forcément dans l’ordre)



Attention, commencer par séparer ce qui là s’appelle Broker en 2 trucs :

-	Un qui écoute et fait les routes API pour déclencher un appel gRPC.
-	Le broker, qui lui se contente de faire du forward.
-	Réfléchir à ce que le broker peut apporter.
-	Ajouter le reflect pour exposer les différents services


-	Le broker, implémenter :

o	Tous les services
o	Celui de tout from ISIN, le faire en streaming pour voir si pertinent
o	Implémenter du channel et du go routine pour qu’il ne se bloque pas
o	Implémenter les contrôles associés, et factoriser autant que possible
o	Mettre un listener d’API, et translater les request http en appel de service gRPC.
o	Implémenter le retour

-	Le Client.  (SecPos retriever)
-	Implémenter PostgreSQL

