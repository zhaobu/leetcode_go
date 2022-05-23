#
# @lc app=leetcode.cn id=151 lang=python3
#
# [151] 颠倒字符串中的单词
#


# @lc code=start
class Solution(object):
    # 去除多余空格
    def trimSpace(self, s: str) -> list:
        # 去除两边空格

        i, j = 0, len(s) - 1
        while i <= j and (s[i] == ' ' or s[j] == ' '):
            if s[i] == ' ':
                i += 1
            if s[j] == ' ':
                j -= 1
        # 去除中间多余空格
        strList = []
        while i <= j:
            if s[i] != ' ':
                strList.append(s[i])
            elif strList[-1] != ' ':
                strList.append(' ')
            i += 1

        return strList

    # 翻转整个字符串
    def reverseStrs(self, strList: list, i: int, j: int) -> None:
        while i < j:
            strList[i], strList[j] = strList[j], strList[i]
            i, j = i + 1, j - 1

    # 翻转每个单词
    def reverseEachWord(self, strList: list) -> str:
        i, j = 0, 1
        while True:
            if j == len(strList) or strList[j] == ' ':
                self.reverseStrs(strList, i, j - 1)
                i = j + 1

            if j == len(strList):
                break
            j += 1

    def reverseWords(self, s: str) -> str:
        strList = self.trimSpace(s)
        # print("s1={}".format(strList))
        self.reverseStrs(strList, 0, len(strList) - 1)
        # print("s1={}".format(strList))
        self.reverseEachWord(strList)
        return "".join(strList)
        # 反转每一个单词


# @lc code=end
