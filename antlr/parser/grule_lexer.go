// Code generated by ANTLR 4.7.
// DO NOT EDIT!

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 39, 409,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4,
	39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44,
	9, 44, 4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9,
	49, 4, 50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 4, 53, 9, 53, 4, 54, 9, 54,
	4, 55, 9, 55, 4, 56, 9, 56, 4, 57, 9, 57, 4, 58, 9, 58, 4, 59, 9, 59, 4,
	60, 9, 60, 4, 61, 9, 61, 4, 62, 9, 62, 4, 63, 9, 63, 4, 64, 9, 64, 4, 65,
	9, 65, 4, 66, 9, 66, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3,
	6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11,
	3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3,
	17, 3, 17, 3, 18, 3, 18, 3, 19, 3, 19, 3, 20, 3, 20, 3, 21, 3, 21, 3, 22,
	3, 22, 3, 23, 3, 23, 3, 24, 3, 24, 3, 25, 3, 25, 3, 26, 3, 26, 3, 27, 3,
	27, 3, 28, 3, 28, 3, 29, 3, 29, 3, 30, 3, 30, 5, 30, 192, 10, 30, 3, 30,
	6, 30, 195, 10, 30, 13, 30, 14, 30, 196, 3, 31, 3, 31, 3, 31, 3, 31, 3,
	31, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33,
	3, 34, 3, 34, 3, 34, 3, 35, 3, 35, 3, 35, 3, 36, 3, 36, 3, 36, 3, 36, 3,
	36, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 38, 3, 38, 3, 38, 3, 38,
	3, 38, 3, 39, 3, 39, 3, 39, 3, 39, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3,
	40, 3, 40, 3, 40, 3, 40, 3, 41, 3, 41, 7, 41, 251, 10, 41, 12, 41, 14,
	41, 254, 11, 41, 3, 42, 3, 42, 3, 42, 3, 42, 6, 42, 260, 10, 42, 13, 42,
	14, 42, 261, 3, 43, 3, 43, 3, 44, 3, 44, 3, 45, 3, 45, 3, 46, 3, 46, 3,
	47, 3, 47, 3, 47, 3, 48, 3, 48, 3, 49, 3, 49, 3, 50, 3, 50, 3, 51, 3, 51,
	3, 51, 3, 52, 3, 52, 3, 52, 3, 53, 3, 53, 3, 53, 3, 54, 3, 54, 3, 55, 3,
	55, 3, 56, 3, 56, 3, 57, 3, 57, 3, 58, 3, 58, 3, 59, 3, 59, 3, 60, 3, 60,
	3, 60, 3, 60, 3, 60, 3, 60, 7, 60, 308, 10, 60, 12, 60, 14, 60, 311, 11,
	60, 3, 60, 3, 60, 3, 61, 3, 61, 3, 61, 3, 61, 3, 61, 3, 61, 7, 61, 321,
	10, 61, 12, 61, 14, 61, 324, 11, 61, 3, 61, 3, 61, 3, 62, 6, 62, 329, 10,
	62, 13, 62, 14, 62, 330, 3, 63, 6, 63, 334, 10, 63, 13, 63, 14, 63, 335,
	5, 63, 338, 10, 63, 3, 63, 3, 63, 6, 63, 342, 10, 63, 13, 63, 14, 63, 343,
	3, 63, 6, 63, 347, 10, 63, 13, 63, 14, 63, 348, 3, 63, 3, 63, 3, 63, 3,
	63, 6, 63, 355, 10, 63, 13, 63, 14, 63, 356, 5, 63, 359, 10, 63, 3, 63,
	3, 63, 6, 63, 363, 10, 63, 13, 63, 14, 63, 364, 3, 63, 3, 63, 3, 63, 6,
	63, 370, 10, 63, 13, 63, 14, 63, 371, 3, 63, 3, 63, 5, 63, 376, 10, 63,
	3, 64, 6, 64, 379, 10, 64, 13, 64, 14, 64, 380, 3, 64, 3, 64, 3, 65, 3,
	65, 3, 65, 3, 65, 7, 65, 389, 10, 65, 12, 65, 14, 65, 392, 11, 65, 3, 65,
	3, 65, 3, 65, 3, 65, 3, 65, 3, 66, 3, 66, 3, 66, 3, 66, 7, 66, 403, 10,
	66, 12, 66, 14, 66, 406, 11, 66, 3, 66, 3, 66, 3, 390, 2, 67, 3, 3, 5,
	2, 7, 2, 9, 2, 11, 2, 13, 2, 15, 2, 17, 2, 19, 2, 21, 2, 23, 2, 25, 2,
	27, 2, 29, 2, 31, 2, 33, 2, 35, 2, 37, 2, 39, 2, 41, 2, 43, 2, 45, 2, 47,
	2, 49, 2, 51, 2, 53, 2, 55, 2, 57, 2, 59, 2, 61, 4, 63, 5, 65, 6, 67, 7,
	69, 8, 71, 9, 73, 10, 75, 11, 77, 12, 79, 13, 81, 14, 83, 15, 85, 16, 87,
	17, 89, 18, 91, 19, 93, 20, 95, 21, 97, 22, 99, 23, 101, 24, 103, 25, 105,
	26, 107, 27, 109, 28, 111, 29, 113, 30, 115, 31, 117, 32, 119, 33, 121,
	34, 123, 35, 125, 36, 127, 37, 129, 38, 131, 39, 3, 2, 35, 3, 2, 50, 59,
	4, 2, 67, 67, 99, 99, 4, 2, 68, 68, 100, 100, 4, 2, 69, 69, 101, 101, 4,
	2, 70, 70, 102, 102, 4, 2, 71, 71, 103, 103, 4, 2, 72, 72, 104, 104, 4,
	2, 73, 73, 105, 105, 4, 2, 74, 74, 106, 106, 4, 2, 75, 75, 107, 107, 4,
	2, 76, 76, 108, 108, 4, 2, 77, 77, 109, 109, 4, 2, 78, 78, 110, 110, 4,
	2, 79, 79, 111, 111, 4, 2, 80, 80, 112, 112, 4, 2, 81, 81, 113, 113, 4,
	2, 82, 82, 114, 114, 4, 2, 83, 83, 115, 115, 4, 2, 84, 84, 116, 116, 4,
	2, 85, 85, 117, 117, 4, 2, 86, 86, 118, 118, 4, 2, 87, 87, 119, 119, 4,
	2, 88, 88, 120, 120, 4, 2, 89, 89, 121, 121, 4, 2, 90, 90, 122, 122, 4,
	2, 91, 91, 123, 123, 4, 2, 92, 92, 124, 124, 4, 2, 67, 92, 99, 124, 5,
	2, 50, 59, 67, 92, 99, 124, 4, 2, 36, 36, 94, 94, 4, 2, 41, 41, 94, 94,
	5, 2, 11, 12, 15, 15, 34, 34, 4, 2, 12, 12, 15, 15, 2, 405, 2, 3, 3, 2,
	2, 2, 2, 61, 3, 2, 2, 2, 2, 63, 3, 2, 2, 2, 2, 65, 3, 2, 2, 2, 2, 67, 3,
	2, 2, 2, 2, 69, 3, 2, 2, 2, 2, 71, 3, 2, 2, 2, 2, 73, 3, 2, 2, 2, 2, 75,
	3, 2, 2, 2, 2, 77, 3, 2, 2, 2, 2, 79, 3, 2, 2, 2, 2, 81, 3, 2, 2, 2, 2,
	83, 3, 2, 2, 2, 2, 85, 3, 2, 2, 2, 2, 87, 3, 2, 2, 2, 2, 89, 3, 2, 2, 2,
	2, 91, 3, 2, 2, 2, 2, 93, 3, 2, 2, 2, 2, 95, 3, 2, 2, 2, 2, 97, 3, 2, 2,
	2, 2, 99, 3, 2, 2, 2, 2, 101, 3, 2, 2, 2, 2, 103, 3, 2, 2, 2, 2, 105, 3,
	2, 2, 2, 2, 107, 3, 2, 2, 2, 2, 109, 3, 2, 2, 2, 2, 111, 3, 2, 2, 2, 2,
	113, 3, 2, 2, 2, 2, 115, 3, 2, 2, 2, 2, 117, 3, 2, 2, 2, 2, 119, 3, 2,
	2, 2, 2, 121, 3, 2, 2, 2, 2, 123, 3, 2, 2, 2, 2, 125, 3, 2, 2, 2, 2, 127,
	3, 2, 2, 2, 2, 129, 3, 2, 2, 2, 2, 131, 3, 2, 2, 2, 3, 133, 3, 2, 2, 2,
	5, 135, 3, 2, 2, 2, 7, 137, 3, 2, 2, 2, 9, 139, 3, 2, 2, 2, 11, 141, 3,
	2, 2, 2, 13, 143, 3, 2, 2, 2, 15, 145, 3, 2, 2, 2, 17, 147, 3, 2, 2, 2,
	19, 149, 3, 2, 2, 2, 21, 151, 3, 2, 2, 2, 23, 153, 3, 2, 2, 2, 25, 155,
	3, 2, 2, 2, 27, 157, 3, 2, 2, 2, 29, 159, 3, 2, 2, 2, 31, 161, 3, 2, 2,
	2, 33, 163, 3, 2, 2, 2, 35, 165, 3, 2, 2, 2, 37, 167, 3, 2, 2, 2, 39, 169,
	3, 2, 2, 2, 41, 171, 3, 2, 2, 2, 43, 173, 3, 2, 2, 2, 45, 175, 3, 2, 2,
	2, 47, 177, 3, 2, 2, 2, 49, 179, 3, 2, 2, 2, 51, 181, 3, 2, 2, 2, 53, 183,
	3, 2, 2, 2, 55, 185, 3, 2, 2, 2, 57, 187, 3, 2, 2, 2, 59, 189, 3, 2, 2,
	2, 61, 198, 3, 2, 2, 2, 63, 203, 3, 2, 2, 2, 65, 208, 3, 2, 2, 2, 67, 213,
	3, 2, 2, 2, 69, 216, 3, 2, 2, 2, 71, 219, 3, 2, 2, 2, 73, 224, 3, 2, 2,
	2, 75, 230, 3, 2, 2, 2, 77, 235, 3, 2, 2, 2, 79, 239, 3, 2, 2, 2, 81, 248,
	3, 2, 2, 2, 83, 255, 3, 2, 2, 2, 85, 263, 3, 2, 2, 2, 87, 265, 3, 2, 2,
	2, 89, 267, 3, 2, 2, 2, 91, 269, 3, 2, 2, 2, 93, 271, 3, 2, 2, 2, 95, 274,
	3, 2, 2, 2, 97, 276, 3, 2, 2, 2, 99, 278, 3, 2, 2, 2, 101, 280, 3, 2, 2,
	2, 103, 283, 3, 2, 2, 2, 105, 286, 3, 2, 2, 2, 107, 289, 3, 2, 2, 2, 109,
	291, 3, 2, 2, 2, 111, 293, 3, 2, 2, 2, 113, 295, 3, 2, 2, 2, 115, 297,
	3, 2, 2, 2, 117, 299, 3, 2, 2, 2, 119, 301, 3, 2, 2, 2, 121, 314, 3, 2,
	2, 2, 123, 328, 3, 2, 2, 2, 125, 375, 3, 2, 2, 2, 127, 378, 3, 2, 2, 2,
	129, 384, 3, 2, 2, 2, 131, 398, 3, 2, 2, 2, 133, 134, 7, 46, 2, 2, 134,
	4, 3, 2, 2, 2, 135, 136, 9, 2, 2, 2, 136, 6, 3, 2, 2, 2, 137, 138, 9, 3,
	2, 2, 138, 8, 3, 2, 2, 2, 139, 140, 9, 4, 2, 2, 140, 10, 3, 2, 2, 2, 141,
	142, 9, 5, 2, 2, 142, 12, 3, 2, 2, 2, 143, 144, 9, 6, 2, 2, 144, 14, 3,
	2, 2, 2, 145, 146, 9, 7, 2, 2, 146, 16, 3, 2, 2, 2, 147, 148, 9, 8, 2,
	2, 148, 18, 3, 2, 2, 2, 149, 150, 9, 9, 2, 2, 150, 20, 3, 2, 2, 2, 151,
	152, 9, 10, 2, 2, 152, 22, 3, 2, 2, 2, 153, 154, 9, 11, 2, 2, 154, 24,
	3, 2, 2, 2, 155, 156, 9, 12, 2, 2, 156, 26, 3, 2, 2, 2, 157, 158, 9, 13,
	2, 2, 158, 28, 3, 2, 2, 2, 159, 160, 9, 14, 2, 2, 160, 30, 3, 2, 2, 2,
	161, 162, 9, 15, 2, 2, 162, 32, 3, 2, 2, 2, 163, 164, 9, 16, 2, 2, 164,
	34, 3, 2, 2, 2, 165, 166, 9, 17, 2, 2, 166, 36, 3, 2, 2, 2, 167, 168, 9,
	18, 2, 2, 168, 38, 3, 2, 2, 2, 169, 170, 9, 19, 2, 2, 170, 40, 3, 2, 2,
	2, 171, 172, 9, 20, 2, 2, 172, 42, 3, 2, 2, 2, 173, 174, 9, 21, 2, 2, 174,
	44, 3, 2, 2, 2, 175, 176, 9, 22, 2, 2, 176, 46, 3, 2, 2, 2, 177, 178, 9,
	23, 2, 2, 178, 48, 3, 2, 2, 2, 179, 180, 9, 24, 2, 2, 180, 50, 3, 2, 2,
	2, 181, 182, 9, 25, 2, 2, 182, 52, 3, 2, 2, 2, 183, 184, 9, 26, 2, 2, 184,
	54, 3, 2, 2, 2, 185, 186, 9, 27, 2, 2, 186, 56, 3, 2, 2, 2, 187, 188, 9,
	28, 2, 2, 188, 58, 3, 2, 2, 2, 189, 191, 7, 71, 2, 2, 190, 192, 7, 47,
	2, 2, 191, 190, 3, 2, 2, 2, 191, 192, 3, 2, 2, 2, 192, 194, 3, 2, 2, 2,
	193, 195, 5, 5, 3, 2, 194, 193, 3, 2, 2, 2, 195, 196, 3, 2, 2, 2, 196,
	194, 3, 2, 2, 2, 196, 197, 3, 2, 2, 2, 197, 60, 3, 2, 2, 2, 198, 199, 5,
	41, 21, 2, 199, 200, 5, 47, 24, 2, 200, 201, 5, 29, 15, 2, 201, 202, 5,
	15, 8, 2, 202, 62, 3, 2, 2, 2, 203, 204, 5, 51, 26, 2, 204, 205, 5, 21,
	11, 2, 205, 206, 5, 15, 8, 2, 206, 207, 5, 33, 17, 2, 207, 64, 3, 2, 2,
	2, 208, 209, 5, 45, 23, 2, 209, 210, 5, 21, 11, 2, 210, 211, 5, 15, 8,
	2, 211, 212, 5, 33, 17, 2, 212, 66, 3, 2, 2, 2, 213, 214, 7, 40, 2, 2,
	214, 215, 7, 40, 2, 2, 215, 68, 3, 2, 2, 2, 216, 217, 7, 126, 2, 2, 217,
	218, 7, 126, 2, 2, 218, 70, 3, 2, 2, 2, 219, 220, 5, 45, 23, 2, 220, 221,
	5, 41, 21, 2, 221, 222, 5, 47, 24, 2, 222, 223, 5, 15, 8, 2, 223, 72, 3,
	2, 2, 2, 224, 225, 5, 17, 9, 2, 225, 226, 5, 7, 4, 2, 226, 227, 5, 29,
	15, 2, 227, 228, 5, 43, 22, 2, 228, 229, 5, 15, 8, 2, 229, 74, 3, 2, 2,
	2, 230, 231, 5, 33, 17, 2, 231, 232, 5, 47, 24, 2, 232, 233, 5, 29, 15,
	2, 233, 234, 5, 29, 15, 2, 234, 76, 3, 2, 2, 2, 235, 236, 5, 33, 17, 2,
	236, 237, 5, 35, 18, 2, 237, 238, 5, 45, 23, 2, 238, 78, 3, 2, 2, 2, 239,
	240, 5, 43, 22, 2, 240, 241, 5, 7, 4, 2, 241, 242, 5, 29, 15, 2, 242, 243,
	5, 23, 12, 2, 243, 244, 5, 15, 8, 2, 244, 245, 5, 33, 17, 2, 245, 246,
	5, 11, 6, 2, 246, 247, 5, 15, 8, 2, 247, 80, 3, 2, 2, 2, 248, 252, 9, 29,
	2, 2, 249, 251, 9, 30, 2, 2, 250, 249, 3, 2, 2, 2, 251, 254, 3, 2, 2, 2,
	252, 250, 3, 2, 2, 2, 252, 253, 3, 2, 2, 2, 253, 82, 3, 2, 2, 2, 254, 252,
	3, 2, 2, 2, 255, 259, 5, 81, 41, 2, 256, 257, 5, 117, 59, 2, 257, 258,
	5, 81, 41, 2, 258, 260, 3, 2, 2, 2, 259, 256, 3, 2, 2, 2, 260, 261, 3,
	2, 2, 2, 261, 259, 3, 2, 2, 2, 261, 262, 3, 2, 2, 2, 262, 84, 3, 2, 2,
	2, 263, 264, 7, 45, 2, 2, 264, 86, 3, 2, 2, 2, 265, 266, 7, 47, 2, 2, 266,
	88, 3, 2, 2, 2, 267, 268, 7, 49, 2, 2, 268, 90, 3, 2, 2, 2, 269, 270, 7,
	44, 2, 2, 270, 92, 3, 2, 2, 2, 271, 272, 7, 63, 2, 2, 272, 273, 7, 63,
	2, 2, 273, 94, 3, 2, 2, 2, 274, 275, 7, 63, 2, 2, 275, 96, 3, 2, 2, 2,
	276, 277, 7, 64, 2, 2, 277, 98, 3, 2, 2, 2, 278, 279, 7, 62, 2, 2, 279,
	100, 3, 2, 2, 2, 280, 281, 7, 64, 2, 2, 281, 282, 7, 63, 2, 2, 282, 102,
	3, 2, 2, 2, 283, 284, 7, 62, 2, 2, 284, 285, 7, 63, 2, 2, 285, 104, 3,
	2, 2, 2, 286, 287, 7, 35, 2, 2, 287, 288, 7, 63, 2, 2, 288, 106, 3, 2,
	2, 2, 289, 290, 7, 61, 2, 2, 290, 108, 3, 2, 2, 2, 291, 292, 7, 125, 2,
	2, 292, 110, 3, 2, 2, 2, 293, 294, 7, 127, 2, 2, 294, 112, 3, 2, 2, 2,
	295, 296, 7, 42, 2, 2, 296, 114, 3, 2, 2, 2, 297, 298, 7, 43, 2, 2, 298,
	116, 3, 2, 2, 2, 299, 300, 7, 48, 2, 2, 300, 118, 3, 2, 2, 2, 301, 309,
	7, 36, 2, 2, 302, 303, 7, 94, 2, 2, 303, 308, 11, 2, 2, 2, 304, 305, 7,
	36, 2, 2, 305, 308, 7, 36, 2, 2, 306, 308, 10, 31, 2, 2, 307, 302, 3, 2,
	2, 2, 307, 304, 3, 2, 2, 2, 307, 306, 3, 2, 2, 2, 308, 311, 3, 2, 2, 2,
	309, 307, 3, 2, 2, 2, 309, 310, 3, 2, 2, 2, 310, 312, 3, 2, 2, 2, 311,
	309, 3, 2, 2, 2, 312, 313, 7, 36, 2, 2, 313, 120, 3, 2, 2, 2, 314, 322,
	7, 41, 2, 2, 315, 316, 7, 94, 2, 2, 316, 321, 11, 2, 2, 2, 317, 318, 7,
	41, 2, 2, 318, 321, 7, 41, 2, 2, 319, 321, 10, 32, 2, 2, 320, 315, 3, 2,
	2, 2, 320, 317, 3, 2, 2, 2, 320, 319, 3, 2, 2, 2, 321, 324, 3, 2, 2, 2,
	322, 320, 3, 2, 2, 2, 322, 323, 3, 2, 2, 2, 323, 325, 3, 2, 2, 2, 324,
	322, 3, 2, 2, 2, 325, 326, 7, 41, 2, 2, 326, 122, 3, 2, 2, 2, 327, 329,
	5, 5, 3, 2, 328, 327, 3, 2, 2, 2, 329, 330, 3, 2, 2, 2, 330, 328, 3, 2,
	2, 2, 330, 331, 3, 2, 2, 2, 331, 124, 3, 2, 2, 2, 332, 334, 5, 5, 3, 2,
	333, 332, 3, 2, 2, 2, 334, 335, 3, 2, 2, 2, 335, 333, 3, 2, 2, 2, 335,
	336, 3, 2, 2, 2, 336, 338, 3, 2, 2, 2, 337, 333, 3, 2, 2, 2, 337, 338,
	3, 2, 2, 2, 338, 339, 3, 2, 2, 2, 339, 341, 7, 48, 2, 2, 340, 342, 5, 5,
	3, 2, 341, 340, 3, 2, 2, 2, 342, 343, 3, 2, 2, 2, 343, 341, 3, 2, 2, 2,
	343, 344, 3, 2, 2, 2, 344, 376, 3, 2, 2, 2, 345, 347, 5, 5, 3, 2, 346,
	345, 3, 2, 2, 2, 347, 348, 3, 2, 2, 2, 348, 346, 3, 2, 2, 2, 348, 349,
	3, 2, 2, 2, 349, 350, 3, 2, 2, 2, 350, 351, 7, 48, 2, 2, 351, 352, 5, 59,
	30, 2, 352, 376, 3, 2, 2, 2, 353, 355, 5, 5, 3, 2, 354, 353, 3, 2, 2, 2,
	355, 356, 3, 2, 2, 2, 356, 354, 3, 2, 2, 2, 356, 357, 3, 2, 2, 2, 357,
	359, 3, 2, 2, 2, 358, 354, 3, 2, 2, 2, 358, 359, 3, 2, 2, 2, 359, 360,
	3, 2, 2, 2, 360, 362, 7, 48, 2, 2, 361, 363, 5, 5, 3, 2, 362, 361, 3, 2,
	2, 2, 363, 364, 3, 2, 2, 2, 364, 362, 3, 2, 2, 2, 364, 365, 3, 2, 2, 2,
	365, 366, 3, 2, 2, 2, 366, 367, 5, 59, 30, 2, 367, 376, 3, 2, 2, 2, 368,
	370, 5, 5, 3, 2, 369, 368, 3, 2, 2, 2, 370, 371, 3, 2, 2, 2, 371, 369,
	3, 2, 2, 2, 371, 372, 3, 2, 2, 2, 372, 373, 3, 2, 2, 2, 373, 374, 5, 59,
	30, 2, 374, 376, 3, 2, 2, 2, 375, 337, 3, 2, 2, 2, 375, 346, 3, 2, 2, 2,
	375, 358, 3, 2, 2, 2, 375, 369, 3, 2, 2, 2, 376, 126, 3, 2, 2, 2, 377,
	379, 9, 33, 2, 2, 378, 377, 3, 2, 2, 2, 379, 380, 3, 2, 2, 2, 380, 378,
	3, 2, 2, 2, 380, 381, 3, 2, 2, 2, 381, 382, 3, 2, 2, 2, 382, 383, 8, 64,
	2, 2, 383, 128, 3, 2, 2, 2, 384, 385, 7, 49, 2, 2, 385, 386, 7, 44, 2,
	2, 386, 390, 3, 2, 2, 2, 387, 389, 11, 2, 2, 2, 388, 387, 3, 2, 2, 2, 389,
	392, 3, 2, 2, 2, 390, 391, 3, 2, 2, 2, 390, 388, 3, 2, 2, 2, 391, 393,
	3, 2, 2, 2, 392, 390, 3, 2, 2, 2, 393, 394, 7, 44, 2, 2, 394, 395, 7, 49,
	2, 2, 395, 396, 3, 2, 2, 2, 396, 397, 8, 65, 3, 2, 397, 130, 3, 2, 2, 2,
	398, 399, 7, 49, 2, 2, 399, 400, 7, 49, 2, 2, 400, 404, 3, 2, 2, 2, 401,
	403, 10, 34, 2, 2, 402, 401, 3, 2, 2, 2, 403, 406, 3, 2, 2, 2, 404, 402,
	3, 2, 2, 2, 404, 405, 3, 2, 2, 2, 405, 407, 3, 2, 2, 2, 406, 404, 3, 2,
	2, 2, 407, 408, 8, 66, 4, 2, 408, 132, 3, 2, 2, 2, 24, 2, 191, 196, 252,
	261, 307, 309, 320, 322, 330, 335, 337, 343, 348, 356, 358, 364, 371, 375,
	380, 390, 404, 5, 3, 64, 2, 3, 65, 3, 3, 66, 4,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "','", "", "", "", "'&&'", "'||'", "", "", "", "", "", "", "", "'+'",
	"'-'", "'/'", "'*'", "'=='", "'='", "'>'", "'<'", "'>='", "'<='", "'!='",
	"';'", "'{'", "'}'", "'('", "')'", "'.'",
}

var lexerSymbolicNames = []string{
	"", "", "RULE", "WHEN", "THEN", "AND", "OR", "TRUE", "FALSE", "NULL_LITERAL",
	"NOT", "SALIENCE", "SIMPLENAME", "DOTTEDNAME", "PLUS", "MINUS", "DIV",
	"MUL", "EQUALS", "ASSIGN", "GT", "LT", "GTE", "LTE", "NOTEQUALS", "SEMICOLON",
	"LR_BRACE", "RR_BRACE", "LR_BRACKET", "RR_BRACKET", "DOT", "DQUOTA_STRING",
	"SQUOTA_STRING", "DECIMAL_LITERAL", "REAL_LITERAL", "SPACE", "COMMENT",
	"LINE_COMMENT",
}

var lexerRuleNames = []string{
	"T__0", "DEC_DIGIT", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J",
	"K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y",
	"Z", "EXPONENT_NUM_PART", "RULE", "WHEN", "THEN", "AND", "OR", "TRUE",
	"FALSE", "NULL_LITERAL", "NOT", "SALIENCE", "SIMPLENAME", "DOTTEDNAME",
	"PLUS", "MINUS", "DIV", "MUL", "EQUALS", "ASSIGN", "GT", "LT", "GTE", "LTE",
	"NOTEQUALS", "SEMICOLON", "LR_BRACE", "RR_BRACE", "LR_BRACKET", "RR_BRACKET",
	"DOT", "DQUOTA_STRING", "SQUOTA_STRING", "DECIMAL_LITERAL", "REAL_LITERAL",
	"SPACE", "COMMENT", "LINE_COMMENT",
}

type gruleLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewgruleLexer(input antlr.CharStream) *gruleLexer {

	l := new(gruleLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "grule.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// gruleLexer tokens.
const (
	gruleLexerT__0            = 1
	gruleLexerRULE            = 2
	gruleLexerWHEN            = 3
	gruleLexerTHEN            = 4
	gruleLexerAND             = 5
	gruleLexerOR              = 6
	gruleLexerTRUE            = 7
	gruleLexerFALSE           = 8
	gruleLexerNULL_LITERAL    = 9
	gruleLexerNOT             = 10
	gruleLexerSALIENCE        = 11
	gruleLexerSIMPLENAME      = 12
	gruleLexerDOTTEDNAME      = 13
	gruleLexerPLUS            = 14
	gruleLexerMINUS           = 15
	gruleLexerDIV             = 16
	gruleLexerMUL             = 17
	gruleLexerEQUALS          = 18
	gruleLexerASSIGN          = 19
	gruleLexerGT              = 20
	gruleLexerLT              = 21
	gruleLexerGTE             = 22
	gruleLexerLTE             = 23
	gruleLexerNOTEQUALS       = 24
	gruleLexerSEMICOLON       = 25
	gruleLexerLR_BRACE        = 26
	gruleLexerRR_BRACE        = 27
	gruleLexerLR_BRACKET      = 28
	gruleLexerRR_BRACKET      = 29
	gruleLexerDOT             = 30
	gruleLexerDQUOTA_STRING   = 31
	gruleLexerSQUOTA_STRING   = 32
	gruleLexerDECIMAL_LITERAL = 33
	gruleLexerREAL_LITERAL    = 34
	gruleLexerSPACE           = 35
	gruleLexerCOMMENT         = 36
	gruleLexerLINE_COMMENT    = 37
)

func (l *gruleLexer) Action(localctx antlr.RuleContext, ruleIndex, actionIndex int) {
	switch ruleIndex {
	case 62:
		l.SPACE_Action(localctx, actionIndex)

	case 63:
		l.COMMENT_Action(localctx, actionIndex)

	case 64:
		l.LINE_COMMENT_Action(localctx, actionIndex)

	default:
		panic("No registered action for: " + fmt.Sprint(ruleIndex))
	}
}

func (l *gruleLexer) SPACE_Action(localctx antlr.RuleContext, actionIndex int) {
	switch actionIndex {
	case 0:
		l.Skip()

	default:
		panic("No registered action for: " + fmt.Sprint(actionIndex))
	}
}
func (l *gruleLexer) COMMENT_Action(localctx antlr.RuleContext, actionIndex int) {
	switch actionIndex {
	case 1:
		l.Skip()

	default:
		panic("No registered action for: " + fmt.Sprint(actionIndex))
	}
}
func (l *gruleLexer) LINE_COMMENT_Action(localctx antlr.RuleContext, actionIndex int) {
	switch actionIndex {
	case 2:
		l.Skip()

	default:
		panic("No registered action for: " + fmt.Sprint(actionIndex))
	}
}
