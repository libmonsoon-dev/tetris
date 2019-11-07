package shape

type Shape []State

var S = Shape{
	{
		{},
		{},
		{o, o, x, x, o, o},
		{o, x, x, o, o, o},
	},
	{
		{},
		{o, o, x, o},
		{o, x, x, o},
		{o, x, o, o},
	},
}

var Z = Shape{
	{
		{},
		{},
		{o, x, x, o, o, o},
		{o, o, x, x, o, o},
	},
	{
		{},
		{o, o, x, o, o, o},
		{o, x, x, o, o, o},
		{o, x, o, o, o, o},
	},
}

var I = Shape{
	{
		{o, o, x, o, o, o},
		{o, o, x, o, o, o},
		{o, o, x, o, o, o},
		{o, o, x, o, o, o},
	},
	{
		{},
		{o, x, x, x, x, o},
	},
}

var O = Shape{
	{
		{},
		{},
		{o, x, x, o},
		{o, x, x, o},
	},
}

var J = Shape{
	{
		{},
		{o, x, o, o, o, o},
		{o, x, x, x, o, o},
	},
	{
		{},
		{o, o, x, x, o, o},
		{o, o, x, o, o, o},
		{o, o, x, o, o, o},
	},
	{
		{},
		{},
		{o, x, x, x, o, o},
		{o, o, o, x, o, o},
	},
	{
		{},
		{o, o, x, o, o},
		{o, o, x, o, o},
		{o, x, x, o, o},
	},
}

var L = Shape{
	{
		{},
		{o, o, o, x, o},
		{o, x, x, x, o},
	},
	{
		{},
		{o, o, x, o, o},
		{o, o, x, o, o},
		{o, o, x, x, o},
	},
	{
		{},
		{},
		{o, x, x, x, o},
		{o, x, o, o, o},
	},
	{
		{},
		{o, x, x, o, o},
		{o, o, x, o, o},
		{o, o, x, o, o},
	},
}

var T = Shape{
	{
		{},
		{o, o, x, o, o},
		{o, x, x, x, o},
	},
	{
		{},
		{o, o, x, o, o},
		{o, o, x, x, o},
		{o, o, x, o, o},
	},
	{
		{},
		{},
		{o, x, x, x, o},
		{o, o, x, o, o},
	},
	{
		{},
		{o, o, x, o, o},
		{o, x, x, o, o},
		{o, o, x, o, o},
	},
}
