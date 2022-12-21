/*
 * Copyright (c) sunist@genius-invokation-simulator-backend, 2022
 * File "reaction.go" LastUpdatedAt 2022/12/12 10:20:12
 */

package definition

// Reaction 反应类型
type Reaction byte

const (
	ReactionNone              Reaction = iota // ReactionNone 无反应
	ReactionMelt                              // ReactionMelt 融化
	ReactionVaporize                          // ReactionVaporize 蒸发
	ReactionOverloaded                        // ReactionOverloaded 超载
	ReactionSuperconduct                      // ReactionSuperconduct 超导
	ReactionFrozen                            // ReactionFrozen 冻结
	ReactionElectroCharged                    // ReactionElectroCharged 感电
	ReactionBurning                           // ReactionBurning 燃烧
	ReactionBloom                             // ReactionBloom 绽放
	ReactionQuicken                           // ReactionQuicken 激化
	ReactionCryoSwirl                         // ReactionCryoSwirl 冰扩散
	ReactionElectroSwirl                      // ReactionElectroSwirl 电扩散
	ReactionHydroSwirl                        // ReactionHydroSwirl 水扩散
	ReactionPyroSwirl                         // ReactionPyroSwirl 火扩散
	ReactionCryoCrystalize                    // ReactionCryoCrystalize 冰结晶
	ReactionElectroCrystalize                 // ReactionElectroCrystalize 雷结晶
	ReactionHydroCrystalize                   // ReactionHydroCrystalize 水结晶
	ReactionPyroCrystalize                    // ReactionPyroCrystalize 火结晶
)