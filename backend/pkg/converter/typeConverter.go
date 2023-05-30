package converter

import (
	"github.com/kioku-project/kioku/pkg/model"
	pbCardDeck "github.com/kioku-project/kioku/services/carddeck/proto"
	pbCollaboration "github.com/kioku-project/kioku/services/collaboration/proto"
	pbUser "github.com/kioku-project/kioku/services/user/proto"
)

func StoreGroupUserRoleToProtoUserIDConverter(role model.GroupUserRole) *pbUser.UserID {
	return &pbUser.UserID{UserID: role.UserID}
}

func StoreGroupAdmissionToProtoUserIDConverter(groupAdmission model.GroupAdmission) *pbUser.UserID {
	return &pbUser.UserID{UserID: groupAdmission.UserID}
}

func StoreGroupAdmissionToProtoGroupInvitationConverter(groupAdmission model.GroupAdmission) *pbCollaboration.GroupInvitation {
	return &pbCollaboration.GroupInvitation{
		AdmissionID: groupAdmission.ID,
		GroupID:     groupAdmission.GroupID,
		GroupName:   groupAdmission.Group.Name,
	}
}

func StoreGroupToProtoGroupConverter(group model.Group) *pbCollaboration.Group {
	return &pbCollaboration.Group{
		GroupID:   group.ID,
		GroupName: group.Name,
		IsDefault: group.IsDefault,
	}
}

func StoreDeckToProtoDeckConverter(deck model.Deck) *pbCardDeck.Deck {
	return &pbCardDeck.Deck{
		DeckID:   deck.ID,
		DeckName: deck.Name,
	}
}

func StoreCardToProtoCardConverter(card model.Card) *pbCardDeck.Card {
	return &pbCardDeck.Card{
		CardID:    card.ID,
		Frontside: card.Frontside,
		Backside:  card.Backside,
	}
}
