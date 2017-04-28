
// generated by gocc; DO NOT EDIT.

package parser

type(
	actionTable [numStates]actionRow
	actionRow struct {
		canRecover bool
		actions [numSymbols]action
	}
)

var actionTab = actionTable{
	actionRow{ // S0
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			shift(2),		/* check */
			nil,		/* service */
			nil,		/* name */
			nil,		/* host */
			nil,		/* expose */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* reload */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* if */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S1
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			accept(true),		/* $ */
			nil,		/* check */
			nil,		/* service */
			nil,		/* name */
			nil,		/* host */
			nil,		/* expose */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* reload */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* if */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S2
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* check */
			shift(3),		/* service */
			nil,		/* name */
			shift(4),		/* host */
			nil,		/* expose */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* reload */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* if */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S3
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* check */
			nil,		/* service */
			shift(5),		/* name */
			nil,		/* host */
			nil,		/* expose */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* reload */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* if */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S4
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* check */
			nil,		/* service */
			nil,		/* name */
			nil,		/* host */
			nil,		/* expose */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* reload */
			nil,		/* alert */
			shift(8),		/* with */
			nil,		/* : */
			nil,		/* ( */
			nil,		/* ) */
			shift(10),		/* if */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S5
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* check */
			nil,		/* service */
			nil,		/* name */
			nil,		/* host */
			shift(14),		/* expose */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* reload */
			nil,		/* alert */
			shift(15),		/* with */
			nil,		/* : */
			nil,		/* ( */
			nil,		/* ) */
			shift(10),		/* if */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S6
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* check */
			nil,		/* service */
			nil,		/* name */
			nil,		/* host */
			nil,		/* expose */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* reload */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
			nil,		/* ( */
			nil,		/* ) */
			shift(10),		/* if */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S7
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(8),		/* $, reduce: Check */
			nil,		/* check */
			nil,		/* service */
			nil,		/* name */
			nil,		/* host */
			nil,		/* expose */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* reload */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
			nil,		/* ( */
			nil,		/* ) */
			shift(10),		/* if */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S8
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* check */
			nil,		/* service */
			shift(18),		/* name */
			nil,		/* host */
			nil,		/* expose */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* reload */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* if */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S9
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(28),		/* $, reduce: RuleList */
			nil,		/* check */
			nil,		/* service */
			nil,		/* name */
			nil,		/* host */
			nil,		/* expose */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* reload */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
			nil,		/* ( */
			nil,		/* ) */
			reduce(28),		/* if, reduce: RuleList */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S10
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* check */
			nil,		/* service */
			shift(20),		/* name */
			nil,		/* host */
			nil,		/* expose */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* reload */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* if */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S11
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* check */
			nil,		/* service */
			nil,		/* name */
			nil,		/* host */
			shift(14),		/* expose */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* reload */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
			nil,		/* ( */
			nil,		/* ) */
			shift(10),		/* if */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S12
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(5),		/* $, reduce: Check */
			nil,		/* check */
			nil,		/* service */
			nil,		/* name */
			nil,		/* host */
			nil,		/* expose */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* reload */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
			nil,		/* ( */
			nil,		/* ) */
			shift(10),		/* if */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S13
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(4),		/* $, reduce: Check */
			nil,		/* check */
			nil,		/* service */
			nil,		/* name */
			nil,		/* host */
			nil,		/* expose */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* reload */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
			nil,		/* ( */
			nil,		/* ) */
			shift(10),		/* if */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S14
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* check */
			nil,		/* service */
			shift(25),		/* name */
			nil,		/* host */
			nil,		/* expose */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* reload */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* if */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S15
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* check */
			nil,		/* service */
			shift(27),		/* name */
			nil,		/* host */
			nil,		/* expose */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* reload */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* if */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S16
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(7),		/* $, reduce: Check */
			nil,		/* check */
			nil,		/* service */
			nil,		/* name */
			nil,		/* host */
			nil,		/* expose */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* reload */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
			nil,		/* ( */
			nil,		/* ) */
			shift(10),		/* if */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S17
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(29),		/* $, reduce: RuleList */
			nil,		/* check */
			nil,		/* service */
			nil,		/* name */
			nil,		/* host */
			nil,		/* expose */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* reload */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
			nil,		/* ( */
			nil,		/* ) */
			reduce(29),		/* if, reduce: RuleList */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S18
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* check */
			nil,		/* service */
			shift(29),		/* name */
			nil,		/* host */
			nil,		/* expose */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* reload */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* if */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S19
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* check */
			nil,		/* service */
			nil,		/* name */
			nil,		/* host */
			nil,		/* expose */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* reload */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
			nil,		/* ( */
			nil,		/* ) */
			reduce(18),		/* if, reduce: ParameterList */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S20
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* check */
			nil,		/* service */
			nil,		/* name */
			nil,		/* host */
			nil,		/* expose */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* reload */
			nil,		/* alert */
			nil,		/* with */
			shift(30),		/* : */
			shift(31),		/* ( */
			nil,		/* ) */
			nil,		/* if */
			reduce(23),		/* operator, reduce: Metric */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S21
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* check */
			nil,		/* service */
			nil,		/* name */
			nil,		/* host */
			nil,		/* expose */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* reload */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* if */
			shift(32),		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S22
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(1),		/* $, reduce: Check */
			nil,		/* check */
			nil,		/* service */
			nil,		/* name */
			nil,		/* host */
			nil,		/* expose */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* reload */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
			nil,		/* ( */
			nil,		/* ) */
			shift(10),		/* if */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S23
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(3),		/* $, reduce: Check */
			nil,		/* check */
			nil,		/* service */
			nil,		/* name */
			nil,		/* host */
			nil,		/* expose */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* reload */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
			nil,		/* ( */
			nil,		/* ) */
			shift(10),		/* if */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S24
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(6),		/* $, reduce: Check */
			nil,		/* check */
			nil,		/* service */
			nil,		/* name */
			nil,		/* host */
			nil,		/* expose */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* reload */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
			nil,		/* ( */
			nil,		/* ) */
			shift(10),		/* if */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S25
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(10),		/* $, reduce: Exposed */
			nil,		/* check */
			nil,		/* service */
			nil,		/* name */
			nil,		/* host */
			nil,		/* expose */
			shift(34),		/* , */
			nil,		/* restart */
			nil,		/* reload */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
			nil,		/* ( */
			nil,		/* ) */
			reduce(10),		/* if, reduce: Exposed */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S26
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(9),		/* $, reduce: ExposedList */
			nil,		/* check */
			nil,		/* service */
			nil,		/* name */
			nil,		/* host */
			nil,		/* expose */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* reload */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
			nil,		/* ( */
			nil,		/* ) */
			reduce(9),		/* if, reduce: ExposedList */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S27
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* check */
			nil,		/* service */
			shift(35),		/* name */
			nil,		/* host */
			nil,		/* expose */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* reload */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* if */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S28
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* check */
			nil,		/* service */
			nil,		/* name */
			nil,		/* host */
			reduce(18),		/* expose, reduce: ParameterList */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* reload */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
			nil,		/* ( */
			nil,		/* ) */
			reduce(18),		/* if, reduce: ParameterList */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S29
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* check */
			nil,		/* service */
			nil,		/* name */
			nil,		/* host */
			nil,		/* expose */
			shift(36),		/* , */
			nil,		/* restart */
			nil,		/* reload */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
			nil,		/* ( */
			nil,		/* ) */
			reduce(20),		/* if, reduce: Parameters */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S30
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* check */
			nil,		/* service */
			shift(37),		/* name */
			nil,		/* host */
			nil,		/* expose */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* reload */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* if */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S31
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* check */
			nil,		/* service */
			shift(38),		/* name */
			nil,		/* host */
			nil,		/* expose */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* reload */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* if */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S32
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* check */
			nil,		/* service */
			shift(39),		/* name */
			nil,		/* host */
			nil,		/* expose */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* reload */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* if */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S33
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(2),		/* $, reduce: Check */
			nil,		/* check */
			nil,		/* service */
			nil,		/* name */
			nil,		/* host */
			nil,		/* expose */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* reload */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
			nil,		/* ( */
			nil,		/* ) */
			shift(10),		/* if */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S34
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* check */
			nil,		/* service */
			shift(25),		/* name */
			nil,		/* host */
			nil,		/* expose */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* reload */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* if */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S35
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* check */
			nil,		/* service */
			nil,		/* name */
			nil,		/* host */
			reduce(20),		/* expose, reduce: Parameters */
			shift(42),		/* , */
			nil,		/* restart */
			nil,		/* reload */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
			nil,		/* ( */
			nil,		/* ) */
			reduce(20),		/* if, reduce: Parameters */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S36
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* check */
			nil,		/* service */
			shift(18),		/* name */
			nil,		/* host */
			nil,		/* expose */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* reload */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* if */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S37
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* check */
			nil,		/* service */
			nil,		/* name */
			nil,		/* host */
			nil,		/* expose */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* reload */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* if */
			reduce(24),		/* operator, reduce: Metric */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S38
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* check */
			nil,		/* service */
			nil,		/* name */
			nil,		/* host */
			nil,		/* expose */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* reload */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
			nil,		/* ( */
			shift(44),		/* ) */
			nil,		/* if */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S39
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* check */
			nil,		/* service */
			nil,		/* name */
			nil,		/* host */
			nil,		/* expose */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* reload */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* if */
			nil,		/* operator */
			reduce(22),		/* then, reduce: HumanAmount */
			reduce(22),		/* for, reduce: HumanAmount */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S40
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* check */
			nil,		/* service */
			nil,		/* name */
			nil,		/* host */
			nil,		/* expose */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* reload */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* if */
			nil,		/* operator */
			shift(45),		/* then */
			shift(46),		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S41
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(11),		/* $, reduce: Exposed */
			nil,		/* check */
			nil,		/* service */
			nil,		/* name */
			nil,		/* host */
			nil,		/* expose */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* reload */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
			nil,		/* ( */
			nil,		/* ) */
			reduce(11),		/* if, reduce: Exposed */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S42
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* check */
			nil,		/* service */
			shift(27),		/* name */
			nil,		/* host */
			nil,		/* expose */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* reload */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* if */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S43
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* check */
			nil,		/* service */
			nil,		/* name */
			nil,		/* host */
			nil,		/* expose */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* reload */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
			nil,		/* ( */
			nil,		/* ) */
			reduce(19),		/* if, reduce: Parameters */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S44
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* check */
			nil,		/* service */
			nil,		/* name */
			nil,		/* host */
			nil,		/* expose */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* reload */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* if */
			reduce(25),		/* operator, reduce: Metric */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S45
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* check */
			nil,		/* service */
			nil,		/* name */
			nil,		/* host */
			nil,		/* expose */
			nil,		/* , */
			shift(50),		/* restart */
			shift(51),		/* reload */
			shift(52),		/* alert */
			nil,		/* with */
			nil,		/* : */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* if */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S46
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* check */
			nil,		/* service */
			shift(53),		/* name */
			nil,		/* host */
			nil,		/* expose */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* reload */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* if */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S47
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* check */
			nil,		/* service */
			nil,		/* name */
			nil,		/* host */
			reduce(19),		/* expose, reduce: Parameters */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* reload */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
			nil,		/* ( */
			nil,		/* ) */
			reduce(19),		/* if, reduce: Parameters */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S48
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(26),		/* $, reduce: Rule */
			nil,		/* check */
			nil,		/* service */
			nil,		/* name */
			nil,		/* host */
			nil,		/* expose */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* reload */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
			nil,		/* ( */
			nil,		/* ) */
			reduce(26),		/* if, reduce: Rule */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S49
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(12),		/* $, reduce: ActionList */
			nil,		/* check */
			nil,		/* service */
			nil,		/* name */
			nil,		/* host */
			nil,		/* expose */
			shift(55),		/* , */
			nil,		/* restart */
			nil,		/* reload */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
			nil,		/* ( */
			nil,		/* ) */
			reduce(12),		/* if, reduce: ActionList */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S50
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(14),		/* $, reduce: Action */
			nil,		/* check */
			nil,		/* service */
			nil,		/* name */
			nil,		/* host */
			nil,		/* expose */
			reduce(14),		/* ,, reduce: Action */
			nil,		/* restart */
			nil,		/* reload */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
			nil,		/* ( */
			nil,		/* ) */
			reduce(14),		/* if, reduce: Action */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S51
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(15),		/* $, reduce: Action */
			nil,		/* check */
			nil,		/* service */
			nil,		/* name */
			nil,		/* host */
			nil,		/* expose */
			reduce(15),		/* ,, reduce: Action */
			nil,		/* restart */
			nil,		/* reload */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
			nil,		/* ( */
			nil,		/* ) */
			reduce(15),		/* if, reduce: Action */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S52
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(16),		/* $, reduce: Action */
			nil,		/* check */
			nil,		/* service */
			shift(56),		/* name */
			nil,		/* host */
			nil,		/* expose */
			reduce(16),		/* ,, reduce: Action */
			nil,		/* restart */
			nil,		/* reload */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
			nil,		/* ( */
			nil,		/* ) */
			reduce(16),		/* if, reduce: Action */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S53
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* check */
			nil,		/* service */
			nil,		/* name */
			nil,		/* host */
			nil,		/* expose */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* reload */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* if */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			reduce(21),		/* cycles, reduce: IntAmount */
			
		},

	},
	actionRow{ // S54
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* check */
			nil,		/* service */
			nil,		/* name */
			nil,		/* host */
			nil,		/* expose */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* reload */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* if */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			shift(57),		/* cycles */
			
		},

	},
	actionRow{ // S55
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* check */
			nil,		/* service */
			nil,		/* name */
			nil,		/* host */
			nil,		/* expose */
			nil,		/* , */
			shift(50),		/* restart */
			shift(51),		/* reload */
			shift(52),		/* alert */
			nil,		/* with */
			nil,		/* : */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* if */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S56
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(17),		/* $, reduce: Action */
			nil,		/* check */
			nil,		/* service */
			nil,		/* name */
			nil,		/* host */
			nil,		/* expose */
			reduce(17),		/* ,, reduce: Action */
			nil,		/* restart */
			nil,		/* reload */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
			nil,		/* ( */
			nil,		/* ) */
			reduce(17),		/* if, reduce: Action */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S57
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* check */
			nil,		/* service */
			nil,		/* name */
			nil,		/* host */
			nil,		/* expose */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* reload */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* if */
			nil,		/* operator */
			shift(59),		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S58
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(13),		/* $, reduce: ActionList */
			nil,		/* check */
			nil,		/* service */
			nil,		/* name */
			nil,		/* host */
			nil,		/* expose */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* reload */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
			nil,		/* ( */
			nil,		/* ) */
			reduce(13),		/* if, reduce: ActionList */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S59
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			nil,		/* $ */
			nil,		/* check */
			nil,		/* service */
			nil,		/* name */
			nil,		/* host */
			nil,		/* expose */
			nil,		/* , */
			shift(50),		/* restart */
			shift(51),		/* reload */
			shift(52),		/* alert */
			nil,		/* with */
			nil,		/* : */
			nil,		/* ( */
			nil,		/* ) */
			nil,		/* if */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	actionRow{ // S60
				canRecover: false,
		actions: [numSymbols]action{
			nil,		/* INVALID */
			reduce(27),		/* $, reduce: Rule */
			nil,		/* check */
			nil,		/* service */
			nil,		/* name */
			nil,		/* host */
			nil,		/* expose */
			nil,		/* , */
			nil,		/* restart */
			nil,		/* reload */
			nil,		/* alert */
			nil,		/* with */
			nil,		/* : */
			nil,		/* ( */
			nil,		/* ) */
			reduce(27),		/* if, reduce: Rule */
			nil,		/* operator */
			nil,		/* then */
			nil,		/* for */
			nil,		/* cycles */
			
		},

	},
	
}

